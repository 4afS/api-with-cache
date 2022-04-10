package service

import (
	"github.com/4afs/api-with-cache/mysql"
	"github.com/google/logger"
)

type Tweet struct {
	User string `json:"user,string"`
	Text string `json:"text,string"`
}

func QueryTimeline(user string) []Tweet {
	db := mysql.Initialize()

	defer db.Close()

	rows, err := db.Query(`
	SELECT user, tweet
	FROM tweets
	WHERE user in (
		SELECT followed
		FROM follows
		WHERE following = ?
		)
	LIMIT 10000
	`, user)

	if err != nil {
		logger.Errorf("query timeline failed: %v", err)
	}

	defer rows.Close()

	tweets := []Tweet{}
	for rows.Next() {
		tweet := Tweet{}
		err := rows.Scan(&tweet.User, &tweet.Text)
		if err != nil {
			logger.Fatalf("tweet scan failed: %v", err)
		}

		tweets = append(tweets, tweet)
	}

	return tweets
}
