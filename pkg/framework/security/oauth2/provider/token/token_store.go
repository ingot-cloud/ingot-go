package token

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"

// Store token 持久化
type Store interface {
	// 根据token读取身份验证信息
	ReadAuthentication(OAuth2AccessToken) (*authentication.OAuth2Authentication, error)
	// 根据token读取身份验证信息
	ReadAuthenticationWith(string) (*authentication.OAuth2Authentication, error)
	// 存储访问令牌
	StoreAccessToken(OAuth2AccessToken, *authentication.OAuth2Authentication)
	// 读取访问令牌
	ReadAccessToken(string) (OAuth2AccessToken, error)
	// 移除访问令牌
	RemoveAccessToken(OAuth2AccessToken)
	// 存储刷新令牌
	StoreRefreshToken(OAuth2RefreshToken, *authentication.OAuth2Authentication)
	// 读取刷新令牌
	ReadRefreshToken(string) (OAuth2RefreshToken, error)
	// 通过刷新令牌读取身份验证信息
	ReadAuthenticationForRefreshToken(OAuth2RefreshToken) (*authentication.OAuth2Authentication, error)
	// 移除刷新令牌
	RemoveRefreshToken(OAuth2RefreshToken)
	// 通过刷新令牌移除访问令牌
	RemoveAccessTokenUsingRefreshToken(OAuth2RefreshToken)
	// 通过身份验证信息获取访问令牌
	GetAccessToken(*authentication.OAuth2Authentication) (OAuth2AccessToken, error)
	// 通过clientID和用户名获取所有访问令牌
	FindTokensByClientIDAndUserName(string, string) ([]OAuth2AccessToken, error)
	// 通过clientID获取所有访问令牌
	FindTokensByClientID(string) ([]OAuth2AccessToken, error)
}
