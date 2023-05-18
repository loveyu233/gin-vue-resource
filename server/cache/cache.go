package cache

import (
	"context"
	"server/global"
	"time"
)

var content = context.Background()

func Get(key string) string {
	return global.RedisClient.Get(content, key).Val()
}

func Set(key string, value interface{}, expiration time.Duration) {
	global.RedisClient.Set(content, key, value, expiration)
}

func Del(key string) {
	global.RedisClient.Del(content, key)
}

func Incr(key string) {
	global.RedisClient.Incr(content, key)
}
