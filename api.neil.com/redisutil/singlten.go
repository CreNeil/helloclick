package redisutil

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var Client *redis.Pool

func init() {
	// 建立连接池
	Client = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
func GetInstance() redis.Conn {
	return Client.Get()
}
