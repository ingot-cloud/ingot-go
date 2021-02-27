package container

import (
	"context"

	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// Instance 容器实例
var Instance *Container

// Container for app
type Container struct {
	HTTPConfig                   config.HTTPConfig
	HTTPConfigurer               config.HTTPConfigurer
	SecurityContainer            *container.SecurityContainer
	OAuth2Container              *container.OAuth2Container
	ResourceServerContainer      *container.ResourceServerContainer
	AuthorizationServerContainer *container.AuthorizationServerContainer
}

// Factory 容器工厂
type Factory func(context.Context) (*Container, func(), error)

// HTTPInjector 注入参数
type HTTPInjector interface {
	GetHTTPConfig() config.HTTPConfig
	GetHTTPConfigurer() config.HTTPConfigurer
}

// BuildContainer 构建 Container
func BuildContainer(httpInjector HTTPInjector, securityContainer container.SecurityAllContainer) *Container {
	return &Container{
		HTTPConfig:                   httpInjector.GetHTTPConfig(),
		HTTPConfigurer:               httpInjector.GetHTTPConfigurer(),
		SecurityContainer:            securityContainer.GetSecurityContainer(),
		OAuth2Container:              securityContainer.GetOAuth2Container(),
		ResourceServerContainer:      securityContainer.GetResourceServerContainer(),
		AuthorizationServerContainer: securityContainer.GetAuthorizationServerContainer(),
	}
}
