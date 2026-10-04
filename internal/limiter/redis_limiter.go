package limiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiter interface {
	Allow(
		ctx context.Context,
		key string,
		max int,
	) (bool, error)
}

type RedisLimiter struct {
	client *redis.Client
}

func NewRedisLimiter(
	client *redis.Client,
) *RedisLimiter {
	return &RedisLimiter{
		client: client,
	}
}

func (r *RedisLimiter) Allow(
	ctx context.Context,
	key string,
	max int,
) (bool, error) {

	redisKey := "rate:v1:" + key

	count, err := r.client.Incr(
		ctx,
		redisKey,
	).Result()

	if err != nil {
		return false, err
	}

	if count == 1 {
		err := r.client.Expire(
			ctx,
			redisKey,
			time.Minute,
		).Err()

		if err != nil {
			return false, err
		}
	}

	return count <= int64(max), nil
}
