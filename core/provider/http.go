package provider

import (
	"ingot/config"
	"ingot/core"
	"ingot/core/middleware"

	"github.com/gin-gonic/gin"
)

// BuildHTTPHandler to get gin.Engine
func BuildHTTPHandler(r core.IRouter) *gin.Engine {
	cfg := config.CONFIG.Server
	gin.SetMode(cfg.Mode)

	app := gin.New()

	app.NoMethod(middleware.NoMethodHandler())

	app.NoRoute(middleware.NoRouteHandler())

	app.Use(middleware.RecoveryMiddleware())

	r.Register(app)

	return app
}
