package container

import (
	bootContainer "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
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
