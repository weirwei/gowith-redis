package redis

import "time"

func (r *Redis) Del(keys ...string) (int64, error) {
	cmd := r.Client.Del(keys...)
	return cmd.Val(), cmd.Err()
}

func (r *Redis) TTL(key string) (time.Duration, error) {
	cmd := r.Client.TTL(key)
	return cmd.Val(), cmd.Err()
}
