package router

import (
	"ingot/api"
	"ingot/config"
	"ingot/core/middleware"
	"ingot/core/security"

	"github.com/gin-gonic/gin"
)

// Router for define api
type Router struct {
	Auth security.Authentication
	Test *api.Test
}

// Register routes
func (r *Router) Register(app *gin.Engine) error {
	cfg := config.CONFIG.Server

	routerGroup := app.Group(cfg.Prefix)

	// authentication
	permitUrls := config.CONFIG.Auth.PermitUrls
	routerGroup.Use(middleware.UserAuthMiddleware(r.Authentication(), middleware.NewPermitWithPrefix(permitUrls...)))

	InitTestRouter(routerGroup, r.Test)
	return nil
}

// Authentication for auth
func (r *Router) Authentication() security.Authentication {
	return r.Auth
}
