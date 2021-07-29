package redis

import (
	"fmt"

	. "github.com/go-redis/redis"
)

func GetRedisClient() *Client {
	redisdb := NewClient(&Options{
		Addr:     "111.230.138.67:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return redisdb
}
