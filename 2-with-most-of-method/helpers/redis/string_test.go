package redis

import (
	"fmt"
	"testing"
)

func TestMSet(t *testing.T) {
	redisdb, err := InitRedisClient()
	if err != nil {
		t.Logf("InitRedisClient err: %v\n", err)
		return
	}
	if err := redisdb.MSet("name", "weirwei", "age", "22"); err != nil {
		t.Logf("MSet err: %v\n", err)
		return
	}

	mGetRes, err := redisdb.MGet("name", "age")
	if err != nil {
		t.Logf("MGet err: %v\n", err)
		return
	}
	t.Logf("%v\n", mGetRes)
	//del, err := redisdb.Del("name", "age")
	//if err != nil {
	//	t.Logf("Del err: %v\n", err)
	//	return
	//}
	//t.Logf("Del success: %d\n", del)
}

func TestMGetPage(t *testing.T) {
	redisdb, err := InitRedisClient()
	if err != nil {
		t.Logf("InitRedisClient err: %v\n", err)
		return
	}
	setData := make([]string, 0, 140)
	getData := make([]string, 0, 70)
	for i := 0; i < 70; i++ {
		setData = append(setData, fmt.Sprintf("key%d", i))
		getData = append(getData, fmt.Sprintf("key%d", i))
		setData = append(setData, fmt.Sprintf("value%d", i))
	}
	if err := redisdb.MSet(setData); err != nil {
		t.Logf("MSet err: %v\n", err)
		return
	}

	mGetRes, err := redisdb.MGet(getData...)
	if err != nil {
		t.Logf("MGet err: %v\n", err)
		return
	}
	t.Logf("%v\n", mGetRes)
	del, err := redisdb.Del(getData...)
	if err != nil {
		t.Logf("Del err: %v\n", err)
		return
	}
	t.Logf("Del success: %d\n", del)

}
