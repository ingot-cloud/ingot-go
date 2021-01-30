package token

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"

// ResourceServerTokenServices 资源服务器 token 服务
type ResourceServerTokenServices interface {
	// 通过token加载身份验证信息
	LoadAuthentication(string) (authentication.OAuth2Authentication, error)
	// 读取指定token详细信息
	ReadAccessToken(string) OAuth2AccessToken
}
