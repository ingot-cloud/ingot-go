package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/preset"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// ResourceServerContainer 资源服务器容器
var ResourceServerContainer = wire.NewSet(wire.Struct(new(container.ResourceServerContainer), "*"))

// ResourceServerContainerFields 资源服务器容器中所有字段
var ResourceServerContainerFields = wire.NewSet(
	ResourceAuthenticationManager,
	ResourceServerWebSecurityConfigurer,
	ResourceServerTokenServices,
	TokenExtractor,
)

// ResourceAuthenticationManager 资源服务器中使用的认证管理器
func ResourceAuthenticationManager(container *container.OAuth2Container, tokenService token.ResourceServerTokenServices, injector container.SecurityInjector) coreAuth.ResourceManager {
	if !injector.EnableResourceServer() {
		return nil
	}
	if injector.GetResourceAuthenticationManager() != nil {
		return injector.GetResourceAuthenticationManager()
	}
	return preset.ResourceAuthenticationManager(container, tokenService)
}

// ResourceServerWebSecurityConfigurer 资源服务器配置
func ResourceServerWebSecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.ResourceManager, injector container.SecurityInjector) security.ResourceServerWebSecurityConfigurer {
	if !injector.EnableResourceServer() {
		return nil
	}
	if injector.GetResourceServerWebSecurityConfigurer() != nil {
		return injector.GetResourceServerWebSecurityConfigurer()
	}
	return preset.ResourceServerWebSecurityConfigurer(tokenExtractor, authenticationManager)
}

// ResourceServerTokenServices 资源服务器 token 服务
func ResourceServerTokenServices(tokenStore token.Store, injector container.SecurityInjector) token.ResourceServerTokenServices {
	if !injector.EnableResourceServer() {
		return nil
	}
	if injector.GetResourceServerTokenServices() != nil {
		return injector.GetResourceServerTokenServices()
	}
	return preset.ResourceServerTokenServices(tokenStore)
}

// TokenExtractor TokenExtrator接口默认实现
func TokenExtractor(injector container.SecurityInjector) authentication.TokenExtractor {
	if !injector.EnableResourceServer() {
		return nil
	}
	if injector.GetTokenExtractor() != nil {
		return injector.GetTokenExtractor()
	}
	return preset.TokenExtractor()
}
