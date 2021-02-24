package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server/middleware"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

func enableDefaultMiddleware(engine *gin.Engine) {
	engine.NoMethod(middleware.NoMethodHandler())
	engine.NoRoute(middleware.NoRouteHandler())
	engine.Use(middleware.RecoveryMiddleware())
}

func enableSecurityMiddleware(engine *gin.Engine, configurers security.WebSecurityConfigurers) {
	config.EnableWebSecurity(engine, configurers)
}
