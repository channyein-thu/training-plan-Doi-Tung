package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatal(" Redis connection failed:", err)
	}

	log.Println(" Redis connected")
	return rdb
}
