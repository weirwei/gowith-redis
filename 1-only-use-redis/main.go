package main

import (
	"fmt"
	"our/helpers/redis"
	"time"
)

func main() {
	redisClient := redis.GetRedisClient()
	if redisClient == nil {
		_ = fmt.Errorf("redisClient is nill")
		return
	}
	name := "weirwei"
	key := "name:weirwei"
	redisClient.Set(key, name, 1*time.Second)
	val := redisClient.Get(key)
	fmt.Printf("%s: %s", key, val)
}
