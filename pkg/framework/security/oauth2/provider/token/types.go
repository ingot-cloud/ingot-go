package token

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/enums"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/request"
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

// ExpiringOAuth2RefreshToken 若实现该接口，代表 RefreshToken 是有过期时间的
type ExpiringOAuth2RefreshToken interface {
	// 获取过期时间
	GetExpiration() time.Time
}

// AuthorizationServerTokenServices 授权服务器
type AuthorizationServerTokenServices interface {
	// 通过身份验证信息创建访问令牌
	CreateAccessToken(*authentication.OAuth2Authentication) (OAuth2AccessToken, error)
	// 通过refresh token和请求信息刷新token
	RefreshAccessToken(string, *request.TokenRequest) (OAuth2AccessToken, error)
	// 根据身份验证信息获取访问令牌
	GetAccessToken(*authentication.OAuth2Authentication) (OAuth2AccessToken, error)
}

// ResourceServerTokenServices 资源服务器 token 服务
type ResourceServerTokenServices interface {
	// 通过access token加载身份验证信息
	LoadAuthentication(string) (*authentication.OAuth2Authentication, error)
	// 读取指定access token详细信息
	ReadAccessToken(string) (OAuth2AccessToken, error)
}

// ConsumerTokenServices token 消费者服务
type ConsumerTokenServices interface {
	// 撤销令牌
	RevokeToken(string) bool
}

// UserAuthenticationConverter 用户map信息和身份验证信息互相转换接口
type UserAuthenticationConverter interface {
	// 在身份验证信息中提取访问令牌使用的信息
	ConvertUserAuthentication(core.Authentication) (map[string]interface{}, error)
	// 从map中提取身份验证信息
	ExtractAuthentication(map[string]interface{}) (core.Authentication, error)
}

// AccessTokenConverter 访问令牌转换器
type AccessTokenConverter interface {
	// 返回访问令牌映射内容
	ConvertAccessToken(OAuth2AccessToken, *authentication.OAuth2Authentication) (map[string]interface{}, error)
	// 根据token value和映射内容提取访问令牌
	ExtractAccessToken(string, map[string]interface{}) (OAuth2AccessToken, error)
	// 根据token映射信息提取身份验证信息
	ExtractAuthentication(map[string]interface{}) (*authentication.OAuth2Authentication, error)
}

// Enhancer token 增强接口
type Enhancer interface {
	// token 增强
	Enhance(OAuth2AccessToken, *authentication.OAuth2Authentication) (OAuth2AccessToken, error)
}

// Enhancers 增强列表
type Enhancers []Enhancer

// Granter 授予Token接口，根据不同的 grantType 实现不同的处理方式
type Granter interface {
	// Grant
	Grant(grantType string, client clientdetails.ClientDetails, tokenRequest *request.TokenRequest) (OAuth2AccessToken, error)
}

// Granters 授权列表
type Granters []Granter
