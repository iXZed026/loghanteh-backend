package queue

import (
	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue(
	client *redis.Client,
) *RedisQueue {
	return &RedisQueue{
		client: client,
	}
}
