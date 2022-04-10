package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/4afs/api-with-cache/migration"
	"github.com/4afs/api-with-cache/redis"
	"github.com/4afs/api-with-cache/service"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

func init() {
	logger.Init("api-with-cache", true, false, io.Discard)
}

func main() {
	if err := migration.Up(); err != nil {
		logger.Fatal("migration failed")
	}

	redis.Clear()

	r := gin.Default()

	r.GET("/timeline/:user", func(c *gin.Context) {
		user := c.Param("user")

		key := fmt.Sprintf("/timeline/%v", user)

		cached, err := redis.Read(key)
		if err == nil {
			c.JSON(200, cached)
			return
		}

		timeline := service.QueryTimeline(user)
		j, err := json.Marshal(&timeline)
		if err != nil {
			logger.Errorf("marshal failed: %v", err)
		}

		err = redis.Set(key, string(j))
		if err != nil {
			logger.Errorf("cache failed: %v", err)
		}

		c.JSON(200, string(j))
	})

	if err := r.Run(); err != nil {
		logger.Error("run failed")
	}

}

// func main() {

// 	r := gin.Default()

// 	r.GET("/ping", func(c *gin.Context) {
// 		logger.Info("/ping called")
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	if err := r.Run(); err != nil {
// 		logger.Error("run failed")
// 	}
// }
