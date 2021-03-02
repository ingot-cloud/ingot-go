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

// SecurityInjector 自定义注入参数，app端实现
type SecurityInjector interface {
	// SecurityContainer
	GetWebSecurityConfigurer() security.WebSecurityConfigurer
	GetHTTPSecurityConfigurer() security.HTTPSecurityConfigurer
	GetWebSecurityConfigurers() security.WebSecurityConfigurers
	GetProviders() coreAuth.Providers
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
	GetResourceServerTokenServices() token.ResourceServerTokenServices
	GetTokenExtractor() authentication.TokenExtractor

	// AuthorizationServerContainer
	GetAuthorizationAuthenticationManager() coreAuth.Manager
	GetAuthorizationServerTokenServices() token.AuthorizationServerTokenServices
	GetConsumerTokenServices() token.ConsumerTokenServices
	GetTokenEndpoint() *endpoint.TokenEndpoint
	GetTokenEnhancer() token.Enhancer
	GetTokenEnhancers() token.Enhancers
	GetTokenGranters() token.Granters
	GetTokenGranter() token.Granter
	GetPasswordTokenGranter() *granter.PasswordTokenGranter
}

// NilSecurityInjector 空实现
type NilSecurityInjector struct {
}

// GetWebSecurityConfigurer 自定义默认配置
func (*NilSecurityInjector) GetWebSecurityConfigurer() security.WebSecurityConfigurer {
	return nil
}

// GetHTTPSecurityConfigurer 自定义默认配置
func (*NilSecurityInjector) GetHTTPSecurityConfigurer() security.HTTPSecurityConfigurer {
	return nil
}

// GetWebSecurityConfigurers 获取自定义值
func (*NilSecurityInjector) GetWebSecurityConfigurers() security.WebSecurityConfigurers {
	return nil
}

// GetProviders 获取自定义值
func (*NilSecurityInjector) GetProviders() coreAuth.Providers {
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
