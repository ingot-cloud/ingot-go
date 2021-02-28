package provider

import (
	"github.com/google/wire"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// ResourceServerContainerSet 资源服务器容器
var ResourceServerContainerSet = wire.NewSet(wire.Struct(new(container.ResourceServerContainer), "*"))

// ResourceServerTokenServices 资源服务器 token 服务
func ResourceServerTokenServices(container *container.OAuth2Container, injector container.SecurityInjector) token.ResourceServerTokenServices {
	if injector.GetResourceServerTokenServices() != nil {
		return injector.GetResourceServerTokenServices()
	}
	return container.DefaultTokenServices
}

// OAuth2SecurityConfigurer 实例化 OAuth2 安全配置
func OAuth2SecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.Manager) *config.OAuth2SecurityConfigurer {
	return config.NewOAuth2SecurityConfigurer(tokenExtractor, authenticationManager)
}

// TokenExtractor TokenExtrator接口默认实现
func TokenExtractor(injector container.SecurityInjector) authentication.TokenExtractor {
	if injector.GetTokenExtractor() != nil {
		return injector.GetTokenExtractor()
	}
	return authentication.NewBearerTokenExtractor()
}

// ResourceAuthenticationManager 资源服务器中使用的认证管理器
func ResourceAuthenticationManager(container *container.OAuth2Container, tokenService token.ResourceServerTokenServices, injector container.SecurityInjector) coreAuth.Manager {
	if injector.GetResourceAuthenticationManager() != nil {
		return injector.GetResourceAuthenticationManager()
	}
	manager := authentication.NewOAuth2AuthenticationManager(tokenService)
	manager.ResourceID = container.Config.ResourceServer.ResourceID
	return manager
}
