package repository

import (
	"Go_Practice/app/model"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type IRedis interface {
	Set(key string, value interface{}, expiretime int) (err error)
	Get(key string) (value string, err error)
}

// Redis 存取值
type Redis struct{}

var redisSingleton *Redis
var redisOnce sync.Once

// RedisIns 獲得單例對象
func RedisIns() IRedis {
	redisOnce.Do(func() {
		redisSingleton = &Redis{}
	})
	return redisSingleton
}

// Set 存入redis值
func (r *Redis) Set(key string, value interface{}, expiretime int) (err error) {
	RedisPool := model.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	_, err = conn.Do("SET", key, value, "EX", expiretime)
	return
}

// Get 取出redis值
func (r *Redis) Get(key string) (value string, err error) {

	RedisPool := model.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	value, err = redis.String(conn.Do("GET", key))
	return
}
