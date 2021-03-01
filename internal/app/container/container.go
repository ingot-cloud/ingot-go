package container

import (
	bootContainer "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// AppContainer app容器
type AppContainer struct {
	*container.NilSecurityInjector
	*bootContainer.DefaultContainer
}

// --- 自定义安全配置 ---

// GetHTTPSecurityConfigurer 设置默认 HttpSecurityConfigurer
func (a *AppContainer) GetHTTPSecurityConfigurer() security.HTTPSecurityConfigurer {
	// todo 如何依赖 security 中的对象
	return nil
}

// 两步容器编译
// 1. 默认实现生成容器，可以依赖默认安全容器中的所有实例，生成前置容器
// 2. 注入前置容器到并生成boot容器，如果有扩展实例则使用新的实例
