package token

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/enums"
)

// OAuth2AccessToken OAuth2 访问令牌
type OAuth2AccessToken interface {
	// 获取额外信息
	GetAdditionalInformation() map[string]interface{}
	// 获取令牌访问范围
	GetScope() []string
	// 获取刷新令牌
	GetRefreshToken() OAuth2RefreshToken
	// 获取令牌类型
	GetTokenType() enums.TokenType
	// 令牌是否过期
	IsExpired() bool
	// 令牌到期时间
	GetExpiration() time.Time
	// 令牌有效期，单位秒
	GetExpiresIn() int
	// 获取令牌值
	GetValue() string
}

// OAuth2RefreshToken OAuth2 刷新令牌
type OAuth2RefreshToken interface {
	GetRefreshTokenValue() string
}
