package server

import (
	"github.com/gin-gonic/gin"
	bootContainer "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server/middleware"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

func enableDefaultMiddleware(engine *gin.Engine) {
	engine.NoMethod(middleware.NoMethodHandler())
	engine.NoRoute(middleware.NoRouteHandler())
	engine.Use(middleware.RecoveryMiddleware())
}

func enableSecurityMiddleware(engine *gin.Engine, boot bootContainer.Container) {
	enableAuthorization := boot.GetSecurityInjector().EnableAuthorizationServer()
	enableResource := boot.GetSecurityInjector().EnableResourceServer()

	if enableAuthorization {
		ingotRouter := ingot.NewRouter(engine.Group(""))
		oauthConfig := boot.GetAuthorizationServerContainer().TokenEndpointHTTPConfigurer
		oauthConfig.Configure(ingotRouter)
		for _, api := range oauthConfig.GetAPI() {
			api.Apply(ingotRouter)
		}
	}

	webConfigurers := boot.GetSecurityContainer().WebSecurityConfigurers
	if enableResource {
		oauth2Auth := config.NewWebSecurityConfigurerAdapter(nil, boot.GetResourceServerContainer().OAuth2SecurityConfigurer)
		webConfigurers = append(webConfigurers, oauth2Auth)
	}

	config.EnableWebSecurity(engine, webConfigurers)
}
