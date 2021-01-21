package store

import (
	"crypto/tls"

	"github.com/go-redis/redis"
)

// RedisParams for create
type RedisParams struct {
	Address   string
	DB        int
	Password  string
	KeyPrefix string
	SSL       bool
}

// RedisClient struct
type RedisClient struct {
	Cli       *redis.Client
	KeyPrefix string
}

// NewRedisClient for create
func NewRedisClient(params *RedisParams) *RedisClient {
	options := &redis.Options{
		Addr:     params.Address,
		DB:       params.DB,
		Password: params.Password,
	}
	if params.SSL {
		options.TLSConfig = &tls.Config{}
	}
	cli := redis.NewClient(options)

	return &RedisClient{
		Cli:       cli,
		KeyPrefix: params.KeyPrefix,
	}
}
