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

// NilSecurityInjector 空实现
type NilSecurityInjector struct {
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

// GetJwtAccessTokenConverter 自定义值
func (*NilSecurityInjector) GetJwtAccessTokenConverter() *store.JwtAccessTokenConverter {
	return nil
}

// GetResourceAuthenticationManager 获取自定义值
func (*NilSecurityInjector) GetResourceAuthenticationManager() coreAuth.Manager {
	return nil
}

// GetResourceServerConfigurer 自定义
func (*NilSecurityInjector) GetResourceServerConfigurer() security.ResourceServerConfigurer {
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

// GetAuthorizationServerConfigurer 自定义
func (*NilSecurityInjector) GetAuthorizationServerConfigurer() security.AuthorizationServerConfigurer {
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
