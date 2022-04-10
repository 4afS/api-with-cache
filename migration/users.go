package migration

import (
	"database/sql"
	"fmt"

	"github.com/google/logger"
)

func DownUsers(db *sql.DB) error {
	_, err := db.Exec(`drop table if exists users`)
	return err
}

func UpUsers(db *sql.DB) error {
	_, err := db.Exec(
		`create table if not exists users (
			name varchar(128) not null primary key,
			password varchar(128) not null,
			updated_at datetime not null default current_timestamp on update current_timestamp
		)
		`)
	if err != nil {
		logger.Fatalf("create users table failed: %v", err)
	}

	query := GenerateQueryValues(1000, func(i int) string {
		return fmt.Sprintf("('user%v', 'password%v')", i, i)
	})

	_, err = db.Exec(fmt.Sprintf("insert into users(name, password) values %v", query))
	if err != nil {
		logger.Fatalf("insert to users failed: %v", err)
	}

	return nil
}
