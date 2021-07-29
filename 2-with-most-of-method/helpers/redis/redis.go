package redis

import (
	"fmt"

	. "github.com/go-redis/redis"
)

func InitRedisClient() (*Redis, error) {
	addr := "111.230.138.67:6379"
	redisdb := NewClient(&Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	pong, err := redisdb.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("%s %v", pong, err)
	}
	return &Redis{
		redisdb,
		addr,
	}, nil
}

type Redis struct {
	*Client
	RemoteAddr string
}
