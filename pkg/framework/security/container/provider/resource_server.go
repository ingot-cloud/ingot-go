package provider

import (
	"github.com/google/wire"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/preset"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// ResourceServerContainer 资源服务器容器
var ResourceServerContainer = wire.NewSet(wire.Struct(new(container.ResourceServerContainer), "*"))

// ResourceServerContainerFields 资源服务器容器中所有字段
var ResourceServerContainerFields = wire.NewSet(
	ResourceAuthenticationManager,
	ResourceServerTokenServices,
	OAuth2SecurityConfigurer,
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

// ResourceServerTokenServices 资源服务器 token 服务
func ResourceServerTokenServices(container *container.OAuth2Container, injector container.SecurityInjector) token.ResourceServerTokenServices {
	if !injector.EnableResourceServer() {
		return nil
	}
	if injector.GetResourceServerTokenServices() != nil {
		return injector.GetResourceServerTokenServices()
	}
	return preset.ResourceServerTokenServices(container)
}

// OAuth2SecurityConfigurer 实例化 OAuth2 安全配置
func OAuth2SecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.ResourceManager, injector container.SecurityInjector) *config.OAuth2SecurityConfigurer {
	if !injector.EnableResourceServer() {
		return nil
	}
	return preset.OAuth2SecurityConfigurer(tokenExtractor, authenticationManager)
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
