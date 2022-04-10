package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/logger"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func Initialize() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("load .env failed")
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")

	dataSource := fmt.Sprintf("%s:%s@(localhost:3306)/api_with_cache?charset=utf8mb4&interpolateParams=true", user, password)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		logger.Fatal("initialize database failed")
	}

	return db
}
