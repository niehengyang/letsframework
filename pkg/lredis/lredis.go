package lredis

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(opts *redis.Options) *redis.Client {
	rdb := redis.NewClient(opts)
	return rdb
}
