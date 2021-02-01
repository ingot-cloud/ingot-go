package token

// DefaultOAuth2RefreshToken 默认实现
type DefaultOAuth2RefreshToken struct {
	Value string
}

// NewDefaultOAuth2RefreshToken 创建默认 OAuth2RefreshToken
func NewDefaultOAuth2RefreshToken(value string) *DefaultOAuth2RefreshToken {
	return &DefaultOAuth2RefreshToken{
		Value: value,
	}
}

// GetRefreshTokenValue 获取RefreshToken
func (token *DefaultOAuth2RefreshToken) GetRefreshTokenValue() string {
	return token.Value
}
