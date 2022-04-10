package migration

import (
	"database/sql"
	"fmt"

	"github.com/google/logger"
)

func DownTweets(db *sql.DB) error {
	_, err := db.Exec(`drop table if exists tweets`)
	return err
}

func UpTweets(db *sql.DB) error {
	_, err := db.Exec(
		`create table if not exists tweets (
			id int not null primary key auto_increment,
			user varchar(128) not null,
			tweet varchar(512) not null,
			updated_at datetime not null default current_timestamp on update current_timestamp,
			foreign key (user) references users(name)
		)
		`)
	if err != nil {
		logger.Errorf("create tweets table failed: %v", err)
	}

	for userId := 1; userId <= 1000; userId++ {
		query := GenerateQueryValues(
			10000,
			func(i int) string {
				return fmt.Sprintf("('user%v', 'tweet_%v')", userId, i)
			},
		)

		_, err := db.Exec(fmt.Sprintf("insert into tweets(user, tweet) values %v", query))
		if err != nil {
			logger.Fatalf("insert to tweets failed: %v", err)
			break
		}
	}

	return nil
}
