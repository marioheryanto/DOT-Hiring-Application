package config

import "github.com/go-redis/redis/v9"

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:55001",
		Password: "redispw", // no password set
		DB:       0,         // use default DB
	})

	return rdb
}
