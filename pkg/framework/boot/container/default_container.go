package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// DefaultContainer for app
type DefaultContainer struct {
	HTTPConfig         config.HTTPConfig
	HTTPConfigurer     api.HTTPConfigurer
	SecurityInjector   container.SecurityInjector
	SecurityContainer  *container.SecurityContainer
	DebugPrintInjector container.PrintSecurityInjector
}

// GetHTTPConfig default
func (c *DefaultContainer) GetHTTPConfig() config.HTTPConfig {
	return c.HTTPConfig
}

// GetHTTPConfigurer default
func (c *DefaultContainer) GetHTTPConfigurer() api.HTTPConfigurer {
	return c.HTTPConfigurer
}

// GetSecurityInjector 安全注入
func (c *DefaultContainer) GetSecurityInjector() container.SecurityInjector {
	return c.SecurityInjector
}

// GetSecurityContainer default
func (c *DefaultContainer) GetSecurityContainer() *container.SecurityContainer {
	return c.SecurityContainer
}
