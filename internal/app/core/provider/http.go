package provider

import (
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core"
	"github.com/ingot-cloud/ingot-go/internal/app/core/middleware"

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
