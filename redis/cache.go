package redis

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/google/logger"
)

var r = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var ctx = context.Background()

func Clear() {
	if err := r.FlushAll(ctx).Err(); err != nil {
		logger.Fatalf("flush all failed: %v", err)
	}
}

func Set(key string, value string) error {
	return r.Set(ctx, key, value, 0).Err()
}

func Read(key string) (string, error) {
	value, err := r.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %v not found", key)
	}
	if err != nil {
		return "", err
	} else {
		return value, nil
	}
}
