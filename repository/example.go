package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/logger"
)

func Ping() {
	db, err := sql.Open("mysql", "root:root@(localhost:3306)/api_with_cache?charset=utf8mb4&interpolateParams=true")
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query(`select * from user`)

	if err != nil {
		logger.Fatal(err)
	}

	var user string
	var password string

	for rows.Next() {
		err := rows.Scan(&user, &password)
		if err != nil {
			logger.Fatal("scan failed")
		}

		logger.Infof("got (%v, %v)", user, password)
	}
}
