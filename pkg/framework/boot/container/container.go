package container

import (
	"context"

	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// Factory 容器工厂
type Factory func(context.Context) (Container, func(), error)

// BootContainer 构建 Container
var BootContainer = wire.NewSet(
	wire.Struct(new(DefaultContainer), "*"),
	wire.Bind(new(Container), new(*DefaultContainer)),
)

// Container boot 容器
type Container interface {
	GetHTTPConfig() config.HTTPConfig
	GetHTTPConfigurer() config.HTTPConfigurer
	GetSecurityContainer() *container.SecurityContainer
	GetOAuth2Container() *container.OAuth2Container
	GetResourceServerContainer() *container.ResourceServerContainer
	GetAuthorizationServerContainer() *container.AuthorizationServerContainer
}

// DefaultContainer for app
type DefaultContainer struct {
	HTTPConfig                   config.HTTPConfig
	HTTPConfigurer               config.HTTPConfigurer
	SecurityContainer            *container.SecurityContainer
	OAuth2Container              *container.OAuth2Container
	ResourceServerContainer      *container.ResourceServerContainer
	AuthorizationServerContainer *container.AuthorizationServerContainer
}

// GetHTTPConfig default
func (c *DefaultContainer) GetHTTPConfig() config.HTTPConfig {
	return c.HTTPConfig
}

// GetHTTPConfigurer default
func (c *DefaultContainer) GetHTTPConfigurer() config.HTTPConfigurer {
	return c.HTTPConfigurer
}

// GetSecurityContainer default
func (c *DefaultContainer) GetSecurityContainer() *container.SecurityContainer {
	return c.SecurityContainer
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
