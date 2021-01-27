package router

import (
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/middleware"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// Router for define api
type Router struct {
	Auth           security.Authentication
	CasbinEnforcer *casbin.SyncedEnforcer
	AuthAPI        *api.Auth
	HTTPConfig     server.Config
	AuthConfig     config.Auth
}

// Register routes
func (r *Router) Register(app *gin.Engine) error {
	routerGroup := app.Group(r.HTTPConfig.Prefix)

	// authentication
	permitUrls := r.AuthConfig.PermitUrls
	routerGroup.Use(middleware.UserAuthMiddleware(r.Auth, middleware.NewPermitWithPrefix(permitUrls...)))
	routerGroup.Use(middleware.CasbinMiddleware(r.CasbinEnforcer, middleware.NewPermitWithPrefix(permitUrls...)))

	InitAuthRouter(routerGroup, r.AuthAPI)
	return nil
}

// Authentication for auth
func (r *Router) Authentication() security.Authentication {
	return r.Auth
}
