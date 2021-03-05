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
)

// PrintSecurityInjector 打印注入
type PrintSecurityInjector interface{}

// SecurityInjector 自定义注入参数，app端实现
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

	// ResourceServerContainer
	GetResourceAuthenticationManager() coreAuth.Manager
	GetResourceServerWebSecurityConfigurer() security.ResourceServerWebSecurityConfigurer
	GetResourceServerTokenServices() token.ResourceServerTokenServices
	GetTokenExtractor() authentication.TokenExtractor

	// AuthorizationServerContainer
	GetAuthorizationAuthenticationManager() coreAuth.Manager
	GetAuthorizationServerWebSecurityConfigurer() security.AuthorizationServerWebSecurityConfigurer
	GetAuthorizationServerTokenServices() token.AuthorizationServerTokenServices
	GetConsumerTokenServices() token.ConsumerTokenServices
	GetTokenEndpoint() *endpoint.TokenEndpoint
	GetTokenEndpointHTTPConfigurer() endpoint.OAuth2HTTPConfigurer
	GetTokenEnhancer() token.Enhancer
	GetTokenEnhancers() token.Enhancers
	GetTokenGranters() token.Granters
	GetTokenGranter() token.Granter
	GetPasswordTokenGranter() *granter.PasswordTokenGranter

	// AuthProvidersContainer
	GetProviders() coreAuth.Providers
}

// NilSecurityInjector 空实现
type NilSecurityInjector struct {
}

// EnableAuthorizationServer 是否开启授权服务
func (*NilSecurityInjector) EnableAuthorizationServer() bool {
	return false
}

// EnableResourceServer 是否开启资源服务
func (*NilSecurityInjector) EnableResourceServer() bool {
	return false
}

// GetWebSecurityConfigurers 获取自定义值
func (*NilSecurityInjector) GetWebSecurityConfigurers() security.WebSecurityConfigurers {
	return nil
}

// GetPasswordEncoder 获取自定义值
func (*NilSecurityInjector) GetPasswordEncoder() password.Encoder {
	return nil
}

// GetUserCache 获取自定义值
func (*NilSecurityInjector) GetUserCache() userdetails.UserCache {
	return nil
}

// GetPreChecker 获取自定义值
func (*NilSecurityInjector) GetPreChecker() userdetails.PreChecker {
	return nil
}

// GetPostChecker 获取自定义值
func (*NilSecurityInjector) GetPostChecker() userdetails.PostChecker {
	return nil
}

// GetUserDetailsService 获取自定义值
func (*NilSecurityInjector) GetUserDetailsService() userdetails.Service {
	return nil
}

// GetClientDetailsService 获取自定义值
func (*NilSecurityInjector) GetClientDetailsService() clientdetails.Service {
	return nil
}

// GetTokenStore 获取自定义值
func (*NilSecurityInjector) GetTokenStore() token.Store {
	return nil
}

// GetAccessTokenConverter 获取自定义值
func (*NilSecurityInjector) GetAccessTokenConverter() token.AccessTokenConverter {
	return nil
}

// GetUserAuthenticationConverter 获取自定义值
func (*NilSecurityInjector) GetUserAuthenticationConverter() token.UserAuthenticationConverter {
	return nil
}

// GetResourceAuthenticationManager 获取自定义值
func (*NilSecurityInjector) GetResourceAuthenticationManager() coreAuth.Manager {
	return nil
}

// GetResourceServerWebSecurityConfigurer 自定义
func (*NilSecurityInjector) GetResourceServerWebSecurityConfigurer() security.ResourceServerWebSecurityConfigurer {
	return nil
}

// GetResourceServerTokenServices 获取自定义值
func (*NilSecurityInjector) GetResourceServerTokenServices() token.ResourceServerTokenServices {
	return nil
}

// GetTokenExtractor 获取自定义值
func (*NilSecurityInjector) GetTokenExtractor() authentication.TokenExtractor {
	return nil
}

// GetAuthorizationAuthenticationManager 获取自定义值
func (*NilSecurityInjector) GetAuthorizationAuthenticationManager() coreAuth.Manager {
	return nil
}

// GetAuthorizationServerWebSecurityConfigurer 自定义
func (*NilSecurityInjector) GetAuthorizationServerWebSecurityConfigurer() security.AuthorizationServerWebSecurityConfigurer {
	return nil
}

// GetAuthorizationServerTokenServices 获取自定义值
func (*NilSecurityInjector) GetAuthorizationServerTokenServices() token.AuthorizationServerTokenServices {
	return nil
}

// GetConsumerTokenServices 获取自定义值
func (*NilSecurityInjector) GetConsumerTokenServices() token.ConsumerTokenServices {
	return nil
}

// GetTokenEndpoint 获取自定义值
func (*NilSecurityInjector) GetTokenEndpoint() *endpoint.TokenEndpoint {
	return nil
}

// GetTokenEndpointHTTPConfigurer 获取自定义值
func (*NilSecurityInjector) GetTokenEndpointHTTPConfigurer() endpoint.OAuth2HTTPConfigurer {
	return nil
}

// GetTokenEnhancer 获取自定义值
func (*NilSecurityInjector) GetTokenEnhancer() token.Enhancer {
	return nil
}

// GetTokenEnhancers 获取自定义值
func (*NilSecurityInjector) GetTokenEnhancers() token.Enhancers {
	return nil
}

// GetTokenGranters 获取自定义值
func (*NilSecurityInjector) GetTokenGranters() token.Granters {
	return nil
}

// GetTokenGranter 获取自定义值
func (*NilSecurityInjector) GetTokenGranter() token.Granter {
	return nil
}

// GetPasswordTokenGranter 获取自定义值
func (*NilSecurityInjector) GetPasswordTokenGranter() *granter.PasswordTokenGranter {
	return nil
}

// GetProviders 获取自定义值
func (*NilSecurityInjector) GetProviders() coreAuth.Providers {
	return nil
}
