package container

import (
	appConfig "github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/service"
	appToken "github.com/ingot-cloud/ingot-go/internal/app/core/security/token"
	bootContainer "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

// AppContainer app容器
type AppContainer struct {
	*container.NilSecurityInjector
	*bootContainer.DefaultPre

	// 容器中的实例
	JwtAccessTokenConverter *store.JwtAccessTokenConverter

	// app中的实例
	SecurityConfig                   appConfig.Security
	ClientDetailsService             *service.ClientDetails
	UserDetailsService               *service.UserDetails
	ResourceServerAdapter            *config.ResourceServerAdapter
	IngotEnhancer                    *appToken.IngotEnhancer
	IngotUserAuthenticationConverter *appToken.IngotUserAuthenticationConverter
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
	return a.ResourceServerAdapter
}

// GetUserDetailsService 获取自定义值
func (a *AppContainer) GetUserDetailsService() userdetails.Service {
	return a.UserDetailsService
}

// GetClientDetailsService 获取自定义值
func (a *AppContainer) GetClientDetailsService() clientdetails.Service {
	return a.ClientDetailsService
}

// GetUserAuthenticationConverter 自定义
func (a *AppContainer) GetUserAuthenticationConverter() token.UserAuthenticationConverter {
	return a.IngotUserAuthenticationConverter
}

// GetTokenEnhancer 自定义token增强
func (a *AppContainer) GetTokenEnhancer() token.Enhancer {
	chain := token.NewEnhancerChain()
	var enhancers []token.Enhancer
	enhancers = append(enhancers, a.IngotEnhancer)
	// 默认追加 jwt enhancer
	enhancers = append(enhancers, a.JwtAccessTokenConverter)
	chain.SetTokenEnhancers(enhancers)
	return chain
}
