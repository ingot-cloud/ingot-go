package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// DefaultPre 为了执行 BuildContainerInjector
type DefaultPre struct {
	HTTPConfig                   config.HTTPConfig
	HTTPConfigurer               api.HTTPConfigurer
	SecurityContainer            *container.SecurityContainer
	OAuth2Container              *container.OAuth2Container
	ResourceServerContainer      *container.ResourceServerContainer
	AuthorizationServerContainer *container.AuthorizationServerContainer
}

// GetHTTPConfig default
func (c *DefaultPre) GetHTTPConfig() config.HTTPConfig {
	return c.HTTPConfig
}

// GetHTTPConfigurer default
func (c *DefaultPre) GetHTTPConfigurer() api.HTTPConfigurer {
	return c.HTTPConfigurer
}

// GetSecurityInjector 安全注入
func (c *DefaultPre) GetSecurityInjector() container.SecurityInjector {
	return nil
}

// GetSecurityContainer default
func (c *DefaultPre) GetSecurityContainer() *container.SecurityContainer {
	return c.SecurityContainer
}

// GetOAuth2Container default
func (c *DefaultPre) GetOAuth2Container() *container.OAuth2Container {
	return c.OAuth2Container
}

// GetResourceServerContainer default
func (c *DefaultPre) GetResourceServerContainer() *container.ResourceServerContainer {
	return c.ResourceServerContainer
}

// GetAuthorizationServerContainer default
func (c *DefaultPre) GetAuthorizationServerContainer() *container.AuthorizationServerContainer {
	return c.AuthorizationServerContainer
}
