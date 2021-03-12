package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// DefaultPre 为了执行 BuildContainerInjector
type DefaultPre struct {
	HTTPConfig        config.HTTPConfig
	HTTPConfigurer    api.HTTPConfigurer
	SecurityContainer container.SecurityContainer
}

// GetHTTPConfig default
func (c *DefaultPre) GetHTTPConfig() config.HTTPConfig {
	return c.HTTPConfig
}

// GetHTTPConfigurer default
func (c *DefaultPre) GetHTTPConfigurer() api.HTTPConfigurer {
	return c.HTTPConfigurer
}

// GetSecurityContainer default
func (c *DefaultPre) GetSecurityContainer() container.SecurityContainer {
	return c.SecurityContainer
}
