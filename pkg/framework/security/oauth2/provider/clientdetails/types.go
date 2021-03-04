package clientdetails

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// ClientDetails 客户端详情
type ClientDetails interface {
	// 获取ClientID
	GetClientID() string
	// 获取可以访问的资源ID
	GetResourceIDs() []string
	// 验证此客户端是否需要秘钥
	IsSecretRequired() bool
	// 获取客户端秘钥
	GetClientSecret() string
	// 是否需要验证 scope
	IsScoped() bool
	// 获取 scope
	GetScope() []string
	// 客户端的授权类型
	GetAuthorizedGrantTypes() []string
	// 获取授权码模式预定义的重定向uri
	GetRegisteredRedirectURI() []string
	// 获取授予客户端的权限
	GetAuthorities() []core.GrantedAuthority
	// 客户端访问令牌有效时间，单位秒
	GetAccessTokenValiditySeconds() int
	// 客户端刷新令牌有效时间，单位秒
	GetRefreshTokenValiditySeconds() int
	// 指定scope是否需要用户授权批准，如果不需要用户批准则返回ture
	IsAutoApprove(string) bool
	// 客户端的额外附加信息
	GetAdditionalInformation() map[string]interface{}
}

// Service 获取客户端详细信息
type Service interface {
	// 根据 clientID 获取客户端详细信息
	LoadClientByClientID(string) (ClientDetails, error)
}
