package redis

import (
	"RisenIOT/backend/pkg/env"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type MyRedis struct {
	rdb *redis.Client
}

// NewRedis 初始化redis
func NewRedis() *MyRedis {
	var r MyRedis

	RedisHost, _ := env.GetEnv("REDIS_HOST")
	RedisPort, _ := env.GetEnv("REDIS_PORT")
	RedisPassword, _ := env.GetEnv("REDIS_PASSWORD")

	r.rdb = redis.NewClient(&redis.Options{
		Addr:     RedisHost + ":" + RedisPort,
		Password: RedisPassword,
		DB:       0, // 默认DB 0
	})

	return &r
}

// Set 设置值
func (r *MyRedis) Set(key string, value string, expiration int) {
	ctx := context.Background()

	ok, _ := r.rdb.Set(ctx, key, value, time.Duration(expiration)).Result()
	if ok != "OK" {
		fmt.Println("设置值失败:", ok)
	}
}

// Get 获取值
func (r *MyRedis) Get(key string) (string, error) {
	ctx := context.Background()

	// 判断是否存在值
	ok, _ := r.rdb.Do(ctx, "exists", key).Result()
	if ok == 0 {
		return "", fmt.Errorf("key %s 不存在", key)
	}

	// 获取值
	value, _ := r.rdb.Get(ctx, key).Result()

	return value, nil
}
