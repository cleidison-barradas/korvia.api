package redis

import (
	"github.com/redis/go-redis/v9"
)

type Config struct {
	RedisHost string
	Password string
	DB int
}

func NewClient(cfg Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost,
		Password: cfg.Password,
		DB: cfg.DB,
	})
}