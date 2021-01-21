package router

import (
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/middleware"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// Router for define api
type Router struct {
	Auth           security.Authentication
	CasbinEnforcer *casbin.SyncedEnforcer
	AuthAPI        *api.Auth
}

// Register routes
func (r *Router) Register(app *gin.Engine) error {
	cfg := config.CONFIG.Server

	routerGroup := app.Group(cfg.Prefix)

	// authentication
	permitUrls := config.CONFIG.Auth.PermitUrls
	routerGroup.Use(middleware.UserAuthMiddleware(r.Auth, middleware.NewPermitWithPrefix(permitUrls...)))
	routerGroup.Use(middleware.CasbinMiddleware(r.CasbinEnforcer, middleware.NewPermitWithPrefix(permitUrls...)))

	InitAuthRouter(routerGroup, r.AuthAPI)
	return nil
}

// Authentication for auth
func (r *Router) Authentication() security.Authentication {
	return r.Auth
}
