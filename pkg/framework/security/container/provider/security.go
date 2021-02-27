package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails/cache"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/factory"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

// SecurityContainerSet 安全容器
var SecurityContainerSet = wire.NewSet(wire.Struct(new(container.SecurityContainer), "*"))

// Providers 所有认证提供者
func Providers(dao *dao.AuthenticationProvider) coreAuth.Providers {
	var providers coreAuth.Providers
	providers = append(providers, dao)
	return providers
}

// PasswordEncoder encoder
func PasswordEncoder(injectorParams *container.CustomContainer) password.Encoder {
	if injectorParams.PasswordEncoder != nil {
		return injectorParams.PasswordEncoder
	}
	return factory.CreateDelegatingPasswordEncoder()
}

// UserCache 用户缓存
func UserCache() userdetails.UserCache {
	return cache.NewNilUserCache()
}

// PreChecker 前置检查器
func PreChecker() userdetails.PreChecker {
	return dao.NewPreChecker()
}

// PostChecker 后置检查器
func PostChecker() userdetails.PostChecker {
	return dao.NewPostChecker()
}

// WebSecurityConfigurers web 安全配置
func WebSecurityConfigurers(injectorParams *container.CustomContainer) security.WebSecurityConfigurers {
	configurers := injectorParams.WebSecurityConfigurers
	configurers = append(configurers, config.NewWebSecurityConfigurerAdapter(nil, nil))
	return configurers
}

// UserDetailsService 用户详情服务
func UserDetailsService(injectorParams *container.CustomContainer) userdetails.Service {
	return injectorParams.UserDetailsService
}

// ClientDetailsService 客户端详情服务
func ClientDetailsService(injectorParams *container.CustomContainer) clientdetails.Service {
	return injectorParams.ClientDetailsService
}

// DaoAuthenticationProviderSet UsernamePasswordAuthenticationToken 认证提供者
var DaoAuthenticationProviderSet = wire.NewSet(wire.Struct(new(dao.AuthenticationProvider), "*"))
