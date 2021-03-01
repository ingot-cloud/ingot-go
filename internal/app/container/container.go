package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// AppContainer app容器
type AppContainer struct {
	HTTPConfig     config.HTTPConfig
	HTTPConfigurer config.HTTPConfigurer

	*container.NilSecurityInjector
}

// GetHTTPConfig 获取配置
func (a *AppContainer) GetHTTPConfig() config.HTTPConfig {
	return a.HTTPConfig
}

// GetHTTPConfigurer 获取配置
func (a *AppContainer) GetHTTPConfigurer() config.HTTPConfigurer {
	return a.HTTPConfigurer
}

// --- 自定义安全配置 ---

// GetHTTPSecurityConfigurer 设置默认 HttpSecurityConfigurer
func (a *AppContainer) GetHTTPSecurityConfigurer() security.HTTPSecurityConfigurer {
	return nil
}
