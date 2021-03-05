package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/preset"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/endpoint"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/granter"
)

// AuthorizationServerContainer 授权服务器容器
var AuthorizationServerContainer = wire.NewSet(wire.Struct(new(container.AuthorizationServerContainer), "*"))

// AuthorizationServerContainerFields 授权服务器容器所有字段
var AuthorizationServerContainerFields = wire.NewSet(
	AuthorizationAuthenticationManager,
	AuthorizationServerWebSecurityConfigurer,
	AuthorizationServerTokenServices,
	ConsumerTokenServices,
	TokenEndpoint,
	TokenEndpointHTTPConfigurer,
	TokenEnhancer,
	TokenGranters,
	TokenGranter,
	PasswordTokenGranter,
)

// AuthorizationAuthenticationManager 授权服务器中的认证管理器
func AuthorizationAuthenticationManager(providerContainer *container.AuthProvidersContainer, injector container.SecurityInjector) authentication.AuthorizationManager {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetAuthorizationAuthenticationManager() != nil {
		return injector.GetAuthorizationAuthenticationManager()
	}
	return preset.AuthorizationAuthenticationManager(providerContainer)
}

// AuthorizationServerWebSecurityConfigurer 授权服务器配置
func AuthorizationServerWebSecurityConfigurer(manager authentication.AuthorizationManager, injector container.SecurityInjector) security.AuthorizationServerWebSecurityConfigurer {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetAuthorizationServerWebSecurityConfigurer() != nil {
		return injector.GetAuthorizationServerWebSecurityConfigurer()
	}
	return preset.AuthorizationServerWebSecurityConfigurer(manager)
}

// AuthorizationServerTokenServices 授权服务器 token 服务
func AuthorizationServerTokenServices(oauth2Container *container.OAuth2Container, common *container.CommonContainer, enhancer token.Enhancer, manager authentication.AuthorizationManager, injector container.SecurityInjector) token.AuthorizationServerTokenServices {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetAuthorizationServerTokenServices() != nil {
		return injector.GetAuthorizationServerTokenServices()
	}
	return preset.AuthorizationServerTokenServices(oauth2Container, common, enhancer, manager)
}

// ConsumerTokenServices 令牌撤销
func ConsumerTokenServices(oauth2Container *container.OAuth2Container, injector container.SecurityInjector) token.ConsumerTokenServices {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetConsumerTokenServices() != nil {
		return injector.GetConsumerTokenServices()
	}
	return preset.ConsumerTokenServices(oauth2Container)
}

// TokenEndpoint 端点
func TokenEndpoint(granter token.Granter, common *container.CommonContainer, injector container.SecurityInjector) *endpoint.TokenEndpoint {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetTokenEndpoint() != nil {
		return injector.GetTokenEndpoint()
	}
	return preset.TokenEndpoint(granter, common)
}

// TokenEndpointHTTPConfigurer 端点配置
func TokenEndpointHTTPConfigurer(tokenEndpoint *endpoint.TokenEndpoint, injector container.SecurityInjector) endpoint.OAuth2HTTPConfigurer {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetTokenEndpointHTTPConfigurer() != nil {
		return injector.GetTokenEndpointHTTPConfigurer()
	}

	return preset.TokenEndpointHTTPConfigurer(tokenEndpoint)
}

// TokenEnhancer token增强，默认使用增强链
func TokenEnhancer(oauth2Container *container.OAuth2Container, injector container.SecurityInjector) token.Enhancer {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetTokenEnhancer() != nil {
		return injector.GetTokenEnhancer()
	}
	return preset.TokenEnhancer(oauth2Container)
}

// TokenGranters 自定义授权
func TokenGranters(injector container.SecurityInjector) token.Granters {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetTokenGranters() != nil {
		return injector.GetTokenGranters()
	}
	return preset.TokenGranters()
}

// TokenGranter token 授权
func TokenGranter(granters token.Granters, password *granter.PasswordTokenGranter, injector container.SecurityInjector) token.Granter {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetTokenGranter() != nil {
		return injector.GetTokenGranter()
	}
	return preset.TokenGranter(granters, password)
}

// PasswordTokenGranter 密码模式授权
func PasswordTokenGranter(tokenServices token.AuthorizationServerTokenServices, manager authentication.AuthorizationManager, injector container.SecurityInjector) *granter.PasswordTokenGranter {
	if !injector.EnableAuthorizationServer() {
		return nil
	}
	if injector.GetPasswordTokenGranter() != nil {
		return injector.GetPasswordTokenGranter()
	}
	return preset.PasswordTokenGranter(tokenServices, manager)
}
