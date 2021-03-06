package http

import (
	"github.com/casbin/casbin/v2"
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/middleware"
	coreApi "github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
)

// APIConfig http 配置
type APIConfig struct {
	CasbinEnforcer *casbin.SyncedEnforcer
	SecurityConfig config.Security

	TestAPI *api.Test
}

// Configure 应用配置
func (c *APIConfig) Configure(app *ingot.Router) {
	// authentication
	permitUrls := c.SecurityConfig.PermitURLs
	app.Use(middleware.CasbinMiddleware(c.CasbinEnforcer, middleware.NewPermitWithPrefix(permitUrls...)))
}

// GetAPI 获取API
func (c *APIConfig) GetAPI() coreApi.Configurers {
	return coreApi.Configurers{c.TestAPI}
}
