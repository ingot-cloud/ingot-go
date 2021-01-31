package token

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/enums"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
)

// DefaultOAuth2AccessToken 访问令牌默认实现
type DefaultOAuth2AccessToken struct {
	Value                 string
	Expiration            time.Time
	TokenType             enums.TokenType
	RefreshToken          OAuth2RefreshToken
	Scope                 []string
	AdditionalInformation map[string]interface{}
}

// GetAdditionalInformation 获取额外信息
func (token *DefaultOAuth2AccessToken) GetAdditionalInformation() map[string]interface{} {
	return token.AdditionalInformation
}

// GetScope 获取令牌访问范围
func (token *DefaultOAuth2AccessToken) GetScope() []string {
	return token.Scope
}

// GetRefreshToken 获取刷新令牌
func (token *DefaultOAuth2AccessToken) GetRefreshToken() OAuth2RefreshToken {
	return token.RefreshToken
}

// GetTokenType 获取令牌类型
func (token *DefaultOAuth2AccessToken) GetTokenType() enums.TokenType {
	if token.TokenType == "" {
		token.TokenType = enums.BearerToken
	}
	return token.TokenType
}

// IsExpired 令牌是否过期
func (token *DefaultOAuth2AccessToken) IsExpired() bool {
	if utils.TimeIsNil(token.GetExpiration()) {
		return false
	}

	return token.GetExpiration().Before(time.Now())
}

// GetExpiration 令牌到期时间
func (token *DefaultOAuth2AccessToken) GetExpiration() time.Time {
	return token.Expiration
}

// GetExpiresIn 令牌有效期，单位秒
func (token *DefaultOAuth2AccessToken) GetExpiresIn() int {
	if utils.TimeIsNil(token.GetExpiration()) {
		return 0
	}

	return int((token.GetExpiration().UnixNano() - time.Now().UnixNano()) / 1e6)
}

// GetValue 获取令牌值
func (token *DefaultOAuth2AccessToken) GetValue() string {
	return token.Value
}
