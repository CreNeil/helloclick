package sessionutil

import (
	"github.com/garyburd/redigo/redis"
	"helloclick/api.neil.com/redisutil"
)

type Session struct {
	Name string
	TTL  int64
}

type SessionInfo struct {
	FirstName string
	LastName  string
}

func (this *Session) Put(key string, value string) error {
	conn := redisutil.GetInstance()
	defer conn.Close()
	_, err := conn.Do("Set", key, value, "EX", this.TTL)
	return err
}

func (this *Session) Get(key string) (string, error) {
	conn := redisutil.GetInstance()
	defer conn.Close()
	value, err := redis.String(conn.Do("Get", key))
	return value, err
}
