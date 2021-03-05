package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"

	coreUtils "github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
)

// All 所有实例
var All = wire.NewSet(
	CommonContainerFields,
	CommonContainer,
	OAuth2ContainerFields,
	OAuth2Container,
	AuthorizationServerContainerFields,
	AuthorizationServerContainer,
	ResourceServerContainerFields,
	ResourceServerContainer,
	AuthProvidersContainer,
	AuthProvidersContainerFields,
	SecurityContainer,
	PrintInjectInstance,
)

// SecurityContainer 安全容器
var SecurityContainer = wire.NewSet(wire.Struct(new(container.SecurityContainer), "*"))

// PrintInjectInstance 打印注入
func PrintInjectInstance(sc *container.SecurityContainer) container.PrintSecurityInjector {
	log.Debugf("PrintInjectInstance %s", coreUtils.GetType(sc))
	var result struct{}
	return &result
}
