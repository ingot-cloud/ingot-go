package token

import "time"

// DefaultExpiringOAuth2RefreshToken 默认带过期信息的 RefreshToken
type DefaultExpiringOAuth2RefreshToken struct {
	*DefaultOAuth2RefreshToken
	expiration time.Time
}

// NewDefaultExpiringOAuth2RefreshToken 创建默认 ExpiringOAuth2RefreshToken
func NewDefaultExpiringOAuth2RefreshToken(value string, expiration time.Time) *DefaultExpiringOAuth2RefreshToken {
	return &DefaultExpiringOAuth2RefreshToken{
		DefaultOAuth2RefreshToken: &DefaultOAuth2RefreshToken{
			Value: value,
		},
		expiration: expiration,
	}
}

// GetExpiration 获取到期时间
func (token *DefaultExpiringOAuth2RefreshToken) GetExpiration() time.Time {
	return token.expiration
}
