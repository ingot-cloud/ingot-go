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

// SecurityContainer 容器汇总
type SecurityContainer interface {
	GetCommonContainer() *CommonContainer
	GetOAuth2Container() *OAuth2Container
	GetResourceServerContainer() *ResourceServerContainer
	GetAuthorizationServerContainer() *AuthorizationServerContainer
	GetAuthProvidersContainer() *AuthProvidersContainer
}

// SecurityInjector 注入器
type SecurityInjector interface {
	// 是否开启授权服务
	EnableAuthorizationServer() bool
	// 是否开启资源服务
	EnableResourceServer() bool

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
