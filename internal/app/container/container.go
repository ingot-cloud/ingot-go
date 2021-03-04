package container

import (
	appConfig "github.com/ingot-cloud/ingot-go/internal/app/config"
	bootContainer "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// AppContainer app容器
type AppContainer struct {
	*container.NilSecurityInjector
	*bootContainer.DefaultPre

	SecurityConfig appConfig.Security
}

// --- 自定义安全配置 ---

// EnableAuthorizationServer 是否开启授权服务
func (a *AppContainer) EnableAuthorizationServer() bool {
	return a.SecurityConfig.EnableAuthorizationServer
}

// EnableResourceServer 是否开启资源服务
func (a *AppContainer) EnableResourceServer() bool {
	return a.SecurityConfig.EnableResourceServer
}

// GetUserDetailsService 获取自定义值
func (a *AppContainer) GetUserDetailsService() userdetails.Service {
	return nil
}

// GetClientDetailsService 获取自定义值
func (a *AppContainer) GetClientDetailsService() clientdetails.Service {
	return nil
}
