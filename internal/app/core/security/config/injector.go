package config

import (
	appConfig "github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/service"
	appToken "github.com/ingot-cloud/ingot-go/internal/app/core/security/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
)

// IngotContainerInjector 安全注入
type IngotContainerInjector struct {
	*container.DefaultContainerInjector

	// app中的实例
	SecurityConfig                   appConfig.Security
	ClientDetailsService             *service.ClientDetails                     `inject:"true"`
	UserDetailsService               *service.UserDetails                       `inject:"true"`
	ResourceServerAdapter            *ResourceServerAdapter                     `inject:"true"`
	IngotEnhancerChain               *appToken.IngotEnhancerChain               `inject:"true"`
	IngotUserAuthenticationConverter *appToken.IngotUserAuthenticationConverter `inject:"true"`
}
