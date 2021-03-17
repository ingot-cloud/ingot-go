package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/container/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
)

// DefaultContainer 默认容器
type DefaultContainer struct {
	HTTPConfig        config.HTTPConfig
	HTTPConfigurer    api.HTTPConfigurer
	SecurityContainer securityContainer.SecurityContainer
}

// GetHTTPConfig default
func (c *DefaultContainer) GetHTTPConfig() config.HTTPConfig {
	return c.HTTPConfig
}

// GetHTTPConfigurer default
func (c *DefaultContainer) GetHTTPConfigurer() api.HTTPConfigurer {
	return c.HTTPConfigurer
}

// GetSecurityContainer default
func (c *DefaultContainer) GetSecurityContainer() securityContainer.SecurityContainer {
	return c.SecurityContainer
}

// DefaultContainerPre 默认容器 pre
type DefaultContainerPre struct {
	HTTPConfig        config.HTTPConfig
	HTTPConfigurer    api.HTTPConfigurer
	SecurityContainer securityContainer.SecurityContainer `container:"true"`
	ContainerInjector ContainerInjector
}

// GetContainerInjector default
func (c *DefaultContainerPre) GetContainerInjector() ContainerInjector {
	return c.ContainerInjector
}

// GetHTTPConfig default
func (c *DefaultContainerPre) GetHTTPConfig() config.HTTPConfig {
	return c.HTTPConfig
}

// GetHTTPConfigurer default
func (c *DefaultContainerPre) GetHTTPConfigurer() api.HTTPConfigurer {
	return c.HTTPConfigurer
}

// GetSecurityContainer default
func (c *DefaultContainerPre) GetSecurityContainer() securityContainer.SecurityContainer {
	return c.SecurityContainer
}

// DefaultContainerInjector 默认容器注入实现
type DefaultContainerInjector struct {
	SecurityInjector securityContainer.SecurityInjector
}

// GetSecurityInjector 获取安全注入
func (ij *DefaultContainerInjector) GetSecurityInjector() securityContainer.SecurityInjector {
	return ij.SecurityInjector
}
