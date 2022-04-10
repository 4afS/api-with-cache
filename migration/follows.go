package migration

import (
	"database/sql"
	"fmt"

	"github.com/google/logger"
)

func DownFollows(db *sql.DB) error {
	_, err := db.Exec(`drop table if exists follows`)
	return err
}

func UpFollows(db *sql.DB) error {
	_, err := db.Exec(
		`create table if not exists follows (
			following varchar(128) not null,
			followed varchar(128) not null,
			updated_at datetime not null default current_timestamp on update current_timestamp,
			foreign key (following) references users(name),
			foreign key (followed) references users(name)
		)
		`)
	if err != nil {
		logger.Errorf("create follows table failed: %v", err)
	}

	for i := 1; i <= 100; i++ {
		query := GenerateQueryValues(1000, func(n int) string {
			return fmt.Sprintf("('user%v', 'user%v')", i, n)
		})

		_, err = db.Exec(fmt.Sprintf("insert into follows(following, followed) values %v", query))
		if err != nil {
			logger.Fatalf("insert to follows failed: %v", err)
		}
	}

	return nil
}
