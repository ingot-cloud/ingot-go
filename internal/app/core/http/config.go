package http

import (
	"github.com/casbin/casbin/v2"
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/middleware"
	bootConfig "github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
)

// APIConfig http 配置
type APIConfig struct {
	// Auth           security.Authentication
	CasbinEnforcer *casbin.SyncedEnforcer
	HTTPConfig     bootConfig.HTTPConfig
	SecurityConfig config.Security

	AuthAPI *api.Auth
}

// Configure 应用配置
func (c *APIConfig) Configure(app *ingot.Router) {
	// authentication
	permitUrls := c.SecurityConfig.PermitURLs
	app.Use(middleware.UserAuthMiddleware(middleware.NewPermitWithPrefix(permitUrls...)))
	app.Use(middleware.CasbinMiddleware(c.CasbinEnforcer, middleware.NewPermitWithPrefix(permitUrls...)))
}

// GetAPI 获取API
func (c *APIConfig) GetAPI() bootConfig.APIConfigurers {
	return bootConfig.APIConfigurers{c.AuthAPI}
}
