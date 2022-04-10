package migration

import (
	"github.com/4afs/api-with-cache/mysql"

	"github.com/google/logger"
)

var db = mysql.Initialize()

func Up() error {
	if err := Down(); err != nil {
		logger.Fatalf("migrate down failed: %v", err)
	}

	// Up

	if err := UpUsers(db); err != nil {
		logger.Fatalf("migrate users failed: %v", err)
	}
	if err := UpFollows(db); err != nil {
		logger.Fatalf("migrate follows failed: %v", err)
	}
	if err := UpTweets(db); err != nil {
		logger.Fatalf("migrate tweets failed: %v", err)
	}

	return nil
}

func Down() error {
	if err := DownTweets(db); err != nil {
		return err
	}

	if err := DownFollows(db); err != nil {
		return err
	}

	if err := DownUsers(db); err != nil {
		return err
	}

	return nil
}
