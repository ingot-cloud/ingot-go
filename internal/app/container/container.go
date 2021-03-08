package container

import (
	appConfig "github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/service"
	bootContainer "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	securityAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/configurer"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// AppContainer app容器
type AppContainer struct {
	*container.NilSecurityInjector
	*bootContainer.DefaultPre

	SecurityConfig       appConfig.Security
	ClientDetailsService *service.ClientDetails
	UserDetailsService   *service.UserDetails
	TokenExtractor       authentication.TokenExtractor
	ResourceManager      securityAuth.ResourceManager
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

// GetResourceServerConfigurer 自定义资源服务配置
func (a *AppContainer) GetResourceServerConfigurer() security.ResourceServerConfigurer {
	parent := configurer.NewResourceServerConfigurer(a.TokenExtractor, a.ResourceManager)
	return config.NewResourceServerAdapter(parent)
}

// GetUserDetailsService 获取自定义值
func (a *AppContainer) GetUserDetailsService() userdetails.Service {
	return a.UserDetailsService
}

// GetClientDetailsService 获取自定义值
func (a *AppContainer) GetClientDetailsService() clientdetails.Service {
	return a.ClientDetailsService
}
