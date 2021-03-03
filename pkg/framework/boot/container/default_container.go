package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// DefaultContainer for app
type DefaultContainer struct {
	HTTPConfig                   config.HTTPConfig
	HTTPConfigurer               api.HTTPConfigurer
	SecurityInjector             container.SecurityInjector
	Common                       *container.Common
	OAuth2Container              *container.OAuth2Container
	ResourceServerContainer      *container.ResourceServerContainer
	AuthorizationServerContainer *container.AuthorizationServerContainer
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
func (c *DefaultContainer) GetSecurityContainer() *container.Common {
	return c.Common
}

// GetOAuth2Container default
func (c *DefaultContainer) GetOAuth2Container() *container.OAuth2Container {
	return c.OAuth2Container
}

// GetResourceServerContainer default
func (c *DefaultContainer) GetResourceServerContainer() *container.ResourceServerContainer {
	return c.ResourceServerContainer
}

// GetAuthorizationServerContainer default
func (c *DefaultContainer) GetAuthorizationServerContainer() *container.AuthorizationServerContainer {
	return c.AuthorizationServerContainer
}
