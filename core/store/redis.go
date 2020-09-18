package store

import (
	"github.com/go-redis/redis"
)

// RedisParams for create
type RedisParams struct {
	Address   string
	DB        int
	Password  string
	KeyPrefix string
}

// RedisClient struct
type RedisClient struct {
	Cli       *redis.Client
	KeyPrefix string
}

// NewRedisClient for create
func NewRedisClient(params *RedisParams) *RedisClient {
	cli := redis.NewClient(&redis.Options{
		Addr:     params.Address,
		DB:       params.DB,
		Password: params.Password,
	})

	return &RedisClient{
		Cli:       cli,
		KeyPrefix: params.KeyPrefix,
	}
}
