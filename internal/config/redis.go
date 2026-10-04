package config

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectRedis(cfg Config) error {

	db, err := strconv.Atoi(cfg.RedisDB)

	if err != nil {
		return err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return err
	}

	Redis = client

	return nil
}
