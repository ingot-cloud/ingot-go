package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/endpoint"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/granter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

// PrintSecurityInjector 打印注入
type PrintSecurityInjector interface{}

// SecurityContainerProxy 安全容器代理
type SecurityContainerProxy interface {
	GetSecurityContainer() SecurityContainer
	GetSecurityInjector() SecurityInjector
}

// SecurityContainerCombine 结合后的安全容器
// 将 SecurityContainer 中的实例替换为 SecurityInjector 中非nil实例
type SecurityContainerCombine interface {
	SecurityContainer
}

// SecurityContainer 安全容器实例
type SecurityContainer interface {
	GetCommonContainer() *CommonContainer
	GetOAuth2Container() *OAuth2Container
	GetResourceServerContainer() *ResourceServerContainer
	GetAuthorizationServerContainer() *AuthorizationServerContainer
	GetAuthProvidersContainer() *AuthProvidersContainer
}

// SecurityInjector 注入器
type SecurityInjector interface {
	// Common
	GetWebSecurityConfigurers() security.WebSecurityConfigurers
	GetPasswordEncoder() password.Encoder
	GetUserCache() userdetails.UserCache
	GetPreChecker() userdetails.PreChecker
	GetPostChecker() userdetails.PostChecker
	GetUserDetailsService() userdetails.Service
	GetClientDetailsService() clientdetails.Service

	// OAuth2Container
	GetTokenStore() token.Store
	GetAccessTokenConverter() token.AccessTokenConverter
	GetUserAuthenticationConverter() token.UserAuthenticationConverter
	GetJwtAccessTokenConverter() *store.JwtAccessTokenConverter

	// ResourceServerContainer
	GetResourceAuthenticationManager() coreAuth.Manager
	GetResourceServerConfigurer() security.ResourceServerConfigurer
	GetResourceServerTokenServices() token.ResourceServerTokenServices
	GetTokenExtractor() authentication.TokenExtractor

	// AuthorizationServerContainer
	GetAuthorizationAuthenticationManager() coreAuth.Manager
	GetAuthorizationServerConfigurer() security.AuthorizationServerConfigurer
	GetAuthorizationServerTokenServices() token.AuthorizationServerTokenServices
	GetConsumerTokenServices() token.ConsumerTokenServices
	GetTokenEndpoint() *endpoint.TokenEndpoint
	GetTokenEndpointHTTPConfigurer() endpoint.OAuth2HTTPConfigurer
	GetTokenEnhancer() token.Enhancer
	GetTokenGranter() token.Granter
	GetPasswordTokenGranter() *granter.PasswordTokenGranter

	// AuthProvidersContainer
	GetProviders() coreAuth.Providers
}
