package config

import (
	appConfig "github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/service"
	appToken "github.com/ingot-cloud/ingot-go/internal/app/core/security/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// IngotSecurityInjector 安全注入
type IngotSecurityInjector struct {
	*container.NilSecurityInjector

	// app中的实例
	SecurityConfig                   appConfig.Security
	ClientDetailsService             *service.ClientDetails                     `inject:"true"`
	UserDetailsService               *service.UserDetails                       `inject:"true"`
	ResourceServerAdapter            *ResourceServerAdapter                     `inject:"true"`
	IngotEnhancerChain               *appToken.IngotEnhancerChain               `inject:"true"`
	IngotUserAuthenticationConverter *appToken.IngotUserAuthenticationConverter `inject:"true"`
}

// --- 自定义安全配置 ---

// GetResourceServerConfigurer 自定义资源服务配置
func (a *IngotSecurityInjector) GetResourceServerConfigurer() security.ResourceServerConfigurer {
	return a.ResourceServerAdapter
}

// GetUserDetailsService 获取自定义值
func (a *IngotSecurityInjector) GetUserDetailsService() userdetails.Service {
	return a.UserDetailsService
}

// GetClientDetailsService 获取自定义值
func (a *IngotSecurityInjector) GetClientDetailsService() clientdetails.Service {
	return a.ClientDetailsService
}

// GetUserAuthenticationConverter 自定义
func (a *IngotSecurityInjector) GetUserAuthenticationConverter() token.UserAuthenticationConverter {
	return a.IngotUserAuthenticationConverter
}

// GetTokenEnhancer 自定义token增强
func (a *IngotSecurityInjector) GetTokenEnhancer() token.Enhancer {
	return a.IngotEnhancerChain
}
