package jwt

import (
	"context"
	"fmt"
	"ingot/core/store"
	"time"
)

// NewTokenStore for jwt
func NewTokenStore(params *store.RedisParams) *TokenStore {
	client := store.NewRedisClient(params)

	return &TokenStore{
		client: client,
	}
}

// TokenStore for jwt
type TokenStore struct {
	client *store.RedisClient
}

// Store 存储令牌
func (t *TokenStore) Store(ctx context.Context, token string, expiration time.Duration) error {
	cmd := t.client.Cli.Set(t.wrapperKey(token), "ignore", expiration)
	return cmd.Err()
}

// Remove 删除令牌
func (t *TokenStore) Remove(ctx context.Context, token string) (bool, error) {
	cmd := t.client.Cli.Del(t.wrapperKey(token))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Check 检查令牌
func (t *TokenStore) Check(ctx context.Context, token string) (bool, error) {
	cmd := t.client.Cli.Exists(t.wrapperKey(token))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Close 关闭资源
func (t *TokenStore) Close() error {
	return t.client.Cli.Close()
}

func (t *TokenStore) wrapperKey(token string) string {
	return fmt.Sprintf("%s%s", t.client.KeyPrefix, token)
}
