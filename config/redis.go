package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}
	log.Println("Redis connected successfully")
}
