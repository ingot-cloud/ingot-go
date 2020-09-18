package provider

import (
	"ingot/config"
	"ingot/core"
	"ingot/core/middleware"

	"github.com/gin-gonic/gin"
)

// HTTPHandlerProvider to get gin.Engine
func HTTPHandlerProvider(r core.IRouter) *gin.Engine {
	cfg := config.CONFIG.Server
	gin.SetMode(cfg.Mode)

	app := gin.New()

	app.NoMethod(middleware.NoMethodHandler())

	app.NoRoute(middleware.NoRouteHandler())

	app.Use(middleware.RecoveryMiddleware())

	r.Register(app)

	return app
}
