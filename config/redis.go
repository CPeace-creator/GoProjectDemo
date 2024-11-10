package config

import (
	"github.com/go-redis/redis"
	"goDemo/global"
	"log"
)

func initRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	global.RedisDB = redisClient
}
