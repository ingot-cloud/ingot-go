package config

import (
	"github.com/gin-gonic/gin"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/container/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// EnableWebSecurity 开启安全认证
func EnableWebSecurity(enableAuthorization, enableResource bool, securityContainer securityContainer.SecurityContainer, engine *gin.Engine) {
	webConfigurers := securityContainer.GetCommonContainer().WebSecurityConfigurers
	// 开启资源服务，增加 OAuth2 安全配置
	if enableResource {
		log.Info(">>>>>> 开启资源服务器安全配置 <<<<<<")
		webConfigurers.Add(securityContainer.GetResourceServerContainer().ResourceServerConfigurer)
	}

	// 开启授权服务，增加配置
	if enableAuthorization {
		log.Info(">>>>>> 开启授权服务器安全配置 <<<<<<")
		webConfigurers.Add(securityContainer.GetAuthorizationServerContainer().AuthorizationServerConfigurer)
	}

	enableWebSecurity(engine, webConfigurers)

	// 增加端点，需要在设置完 WebSecurity 后在进行开启端点，下面的端点不会执行过滤器链
	if enableAuthorization {
		enableOAuth2Endpoint(securityContainer, engine)
	}

}

// enableOAuth2Endpoint 开启端点
func enableOAuth2Endpoint(securityContainer securityContainer.SecurityContainer, engine *gin.Engine) {
	ingotRouter := ingot.NewRouter(engine.Group(""))
	oauthConfig := securityContainer.GetAuthorizationServerContainer().TokenEndpointHTTPConfigurer
	oauthConfig.Configure(ingotRouter)
	for _, api := range oauthConfig.GetAPI() {
		api.Apply(ingotRouter)
	}
}
