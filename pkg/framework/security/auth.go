package security

import (
	"context"
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// User for Auth
type User struct {
	ID       types.ID `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// AccessToken interface
type AccessToken interface {
	// 获取 Token
	GetValue() string
	// 获取 Token 类型
	GetTokenType() string
	// 到期时间
	GetExpiration() int64
	// 是否过期
	IsExpired() bool
}

// Authentication interface
type Authentication interface {
	// 生成令牌
	GenerateToken(ctx context.Context, user User) (AccessToken, error)

	// 解析用户
	ParseUser(ctx context.Context, accessToken string) (*User, error)

	// 销毁令牌
	RevokeToken(ctx context.Context, accessToken string) error

	// 释放资源
	Release() error

	// 获取TokenStore
	GetStore() TokenStore
}

// TokenStore interface
type TokenStore interface {
	// 存储令牌
	Store(ctx context.Context, token string, expiration time.Duration) error
	// 删除令牌
	Remove(ctx context.Context, token string) (bool, error)
	// 检查令牌
	Check(ctx context.Context, token string) (bool, error)
	// 关闭资源
	Close() error
}
