package config

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
)

// EnableWebSecurity 开启安全认证
func EnableWebSecurity(enableAuthorization, enableResource bool, securityContainer *container.SecurityContainer, engine *gin.Engine) {
	// 增加两个过滤器链路，一个是资源服务器的认证过滤器，一个是授权服务器的链路（包含basicFilter）

	webConfigurers := securityContainer.CommonContainer.WebSecurityConfigurers
	// 开启资源服务，增加 OAuth2 安全配置
	if enableResource {
		webConfigurers.Add(securityContainer.ResourceServerContainer.ResourceServerWebSecurityConfigurer)
	}

	// 开启授权服务，增加配置和token 端点
	if enableAuthorization {
		webConfigurers.Add(securityContainer.AuthorizationServerContainer.AuthorizationServerWebSecurityConfigurer)
		config.EnableOAuth2Endpoint(securityContainer, engine)
	}

	enableWebSecurity(engine, webConfigurers)
}
