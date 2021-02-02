package http

import (
	"github.com/casbin/casbin/v2"
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/middleware"
	bootConfig "github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
)

// APIConfig http 配置
type APIConfig struct {
	Auth           security.Authentication
	CasbinEnforcer *casbin.SyncedEnforcer
	HTTPConfig     bootConfig.HTTPConfig
	AuthConfig     config.Auth

	AuthAPI *api.Auth
}

// Configure 应用配置
func (c *APIConfig) Configure(app *ingot.Router) {
	// authentication
	permitUrls := c.AuthConfig.PermitUrls
	app.Use(middleware.UserAuthMiddleware(c.Auth, middleware.NewPermitWithPrefix(permitUrls...)))
	app.Use(middleware.CasbinMiddleware(c.CasbinEnforcer, middleware.NewPermitWithPrefix(permitUrls...)))

	for _, api := range c.getAPI() {
		api.Apply(app)
	}
}

// 获取API
func (c *APIConfig) getAPI() bootConfig.APIConfigurers {
	return bootConfig.APIConfigurers{c.AuthAPI}
}
