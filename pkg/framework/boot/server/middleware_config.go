package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server/middleware"
	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/config"
)

func enableDefaultMiddleware(engine *gin.Engine) {
	engine.NoMethod(middleware.NoMethodHandler())
	engine.NoRoute(middleware.NoRouteHandler())
	engine.Use(middleware.RecoveryMiddleware())
}

func enableSecurityMiddleware(engine *gin.Engine, boot container.Container) {
	oauth2Config := boot.GetSecurityContainer().GetOAuth2Container().OAuth2Config
	enableAuthorization := oauth2Config.AuthorizationServer.Enable
	enableResource := oauth2Config.ResourceServer.Enable
	config.EnableWebSecurity(enableAuthorization, enableResource, boot.GetSecurityContainer(), engine)
}
