package redis

import "time"

const (
	_CHUNK_SIZE = 32
)

func (r *Redis) Get(key string) (string, error) {
	cmd := r.Client.Get(key)
	return cmd.Val(), cmd.Err()
}

func (r *Redis) MGet(keys ...string) ([]interface{}, error) {
	res := make([]interface{}, 0, len(keys))
	for i := 0; i <= len(keys)/_CHUNK_SIZE; i++ {
		end := 0
		if i < len(keys)/_CHUNK_SIZE {
			end = (i + 1) * _CHUNK_SIZE
		} else {
			end = len(keys)
		}
		chunk := keys[i*_CHUNK_SIZE : end]
		cmd := r.Client.MGet(chunk...)
		if cmd.Err() != nil {
			return nil, cmd.Err()
		}
		res = append(res, cmd.Val()...)
	}
	return res, nil
}

func (r *Redis) Set(key, value string) error {
	return r.SetEx(key, value, Forever)
}

func (r *Redis) SetEx(key, value string, duration time.Duration) error {
	return r.Client.Set(key, value, duration).Err()
}

func (r *Redis) MSet(pairs ...interface{}) error {
	cmd := r.Client.MSet(pairs...)
	return cmd.Err()
}
