package container

import (
	bootContainer "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// AppContainer app容器
type AppContainer struct {
	*container.NilSecurityInjector
	*bootContainer.DefaultContainer

	OAuth2SecurityConfigurer *config.OAuth2SecurityConfigurer
}

// --- 自定义安全配置 ---

// GetHTTPSecurityConfigurer 设置默认 HttpSecurityConfigurer
func (a *AppContainer) GetHTTPSecurityConfigurer() security.HTTPSecurityConfigurer {
	return a.OAuth2SecurityConfigurer
}

// GetUserDetailsService 获取自定义值
func (a *AppContainer) GetUserDetailsService() userdetails.Service {
	return nil
}

// GetClientDetailsService 获取自定义值
func (a *AppContainer) GetClientDetailsService() clientdetails.Service {
	return nil
}
