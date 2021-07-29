package redis

import "testing"

func TestTTL(t *testing.T) {
	redisdb, err := InitRedisClient()
	if err != nil {
		t.Logf("InitRedisClient err: %v\n", err)
		return
	}
	ttl, err := redisdb.TTL("name")
	if err != nil {
		t.Logf("ttl err: %v\n", ttl)
		return
	}
	t.Logf("ttl name: %v\n", ttl)
}
