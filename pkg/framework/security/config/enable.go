package config

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// EnableWebSecurity 开启安全认证
func EnableWebSecurity(enableAuthorization, enableResource bool, securityContainer *container.SecurityContainer, engine *gin.Engine) {
	// 增加两个过滤器链路，一个是资源服务器的认证过滤器，一个是授权服务器的链路（包含basicFilter）

	webConfigurers := securityContainer.CommonContainer.WebSecurityConfigurers
	// 开启资源服务，增加 OAuth2 安全配置
	if enableResource {
		webConfigurers.Add(securityContainer.ResourceServerContainer.ResourceServerWebSecurityConfigurer)
	}

	// 开启授权服务，增加配置
	if enableAuthorization {
		webConfigurers.Add(securityContainer.AuthorizationServerContainer.AuthorizationServerWebSecurityConfigurer)
	}

	enableWebSecurity(engine, webConfigurers)

	// 增加端点，需要在设置完 WebSecurity 后在进行开启端点，下面的端点不会执行过滤器链
	if enableAuthorization {
		enableOAuth2Endpoint(securityContainer, engine)
	}

}

// enableOAuth2Endpoint 开启端点
func enableOAuth2Endpoint(securityContainer *container.SecurityContainer, engine *gin.Engine) {
	ingotRouter := ingot.NewRouter(engine.Group(""))
	oauthConfig := securityContainer.AuthorizationServerContainer.TokenEndpointHTTPConfigurer
	oauthConfig.Configure(ingotRouter)
	for _, api := range oauthConfig.GetAPI() {
		api.Apply(ingotRouter)
	}
}
