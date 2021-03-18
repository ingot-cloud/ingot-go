package container

import (
	"context"
	"reflect"

	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/container/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
)

// Factory 容器工厂
type Factory func(context.Context) (Container, func(), error)

// ContainerPre 前置容器
type ContainerPre interface {
	Container
	GetContainerInjector() ContainerInjector
}

// ContainerPrint 打印容器实例
type ContainerPrint interface {
	Container
}

// Container boot 容器
type Container interface {
	GetHTTPConfig() config.HTTPConfig
	GetHTTPConfigurer() api.HTTPConfigurer
	GetSecurityContainer() securityContainer.SecurityContainer
}

// ContainerInjector 容器注入器
type ContainerInjector interface {
	// 根据类型获取值
	GetValue(ContainerInjector, reflect.Type) reflect.Value
}
