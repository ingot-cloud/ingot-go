package post

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/pre"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

func enableResourceServer(sc container.SecurityContainerCombine) bool {
	oauth2Container := sc.GetOAuth2Container()
	return oauth2Container.OAuth2Config.ResourceServer.Enable
}

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
func ResourceAuthenticationManager(sc container.SecurityContainerCombine) coreAuth.ResourceManager {
	if !enableResourceServer(sc) {
		return nil
	}
	config := sc.GetOAuth2Container().OAuth2Config
	return pre.ResourceAuthenticationManager(config, sc.GetResourceServerContainer().ResourceServerTokenServices)
}

// ResourceServerConfigurer 资源服务器配置
func ResourceServerConfigurer(sc container.SecurityContainerCombine) security.ResourceServerConfigurer {
	if !enableResourceServer(sc) {
		return nil
	}
	return pre.ResourceServerConfigurer(sc.GetResourceServerContainer().TokenExtractor, sc.GetAuthorizationServerContainer().AuthenticationManager)
}

// ResourceServerTokenServices 资源服务器 token 服务
func ResourceServerTokenServices(sc container.SecurityContainerCombine) token.ResourceServerTokenServices {
	if !enableResourceServer(sc) {
		return nil
	}
	return pre.ResourceServerTokenServices(sc.GetOAuth2Container().TokenStore)
}

// TokenExtractor TokenExtrator接口默认实现
func TokenExtractor(sc container.SecurityContainerCombine) authentication.TokenExtractor {
	if !enableResourceServer(sc) {
		return nil
	}
	return pre.TokenExtractor()
}
