package preset

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/configurer"
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
func ResourceAuthenticationManager(container *container.OAuth2Container, tokenService token.ResourceServerTokenServices) coreAuth.ResourceManager {
	manager := authentication.NewOAuth2AuthenticationManager(tokenService)
	manager.ResourceID = container.Config.ResourceServer.ResourceID
	return manager
}

// ResourceServerWebSecurityConfigurer 资源服务器配置
func ResourceServerWebSecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.ResourceManager) security.ResourceServerWebSecurityConfigurer {
	return configurer.NewResourceServerWebSecurityConfigurer(tokenExtractor, authenticationManager)
}

// ResourceServerTokenServices 资源服务器 token 服务
func ResourceServerTokenServices(tokenStore token.Store) token.ResourceServerTokenServices {
	service := token.NewDefaultTokenServices(tokenStore)
	return service
}

// TokenExtractor TokenExtrator接口默认实现
func TokenExtractor() authentication.TokenExtractor {
	return authentication.NewBearerTokenExtractor()
}
