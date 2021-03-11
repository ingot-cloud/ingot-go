package preset

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/configurer"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// ResourceServerContainer 资源服务器容器
var ResourceServerContainer = wire.NewSet(wire.Struct(new(container.ResourceServerContainer), "*"))

// ResourceServerContainerFields 资源服务器容器中所有字段
var ResourceServerContainerFields = wire.NewSet(
	ResourceAuthenticationManager,
	ResourceServerConfigurer,
	ResourceServerTokenServices,
	TokenExtractor,
)

// ResourceAuthenticationManager 资源服务器中使用的认证管理器
func ResourceAuthenticationManager(oauthConfig config.OAuth2, tokenService token.ResourceServerTokenServices) coreAuth.ResourceManager {
	manager := authentication.NewOAuth2AuthenticationManager(tokenService)
	manager.ResourceID = oauthConfig.ResourceServer.ResourceID
	return manager
}

// ResourceServerConfigurer 资源服务器配置
func ResourceServerConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.ResourceManager) security.ResourceServerConfigurer {
	return configurer.NewResourceServerConfigurer(tokenExtractor, authenticationManager)
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
