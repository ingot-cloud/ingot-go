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

// SecurityContainer 安全容器
var SecurityContainer = wire.NewSet(wire.Struct(new(container.SecurityContainer), "*"))

// SecurityContainerFields 安全容器所有字段
var SecurityContainerFields = wire.NewSet(
	Providers,
	PasswordEncoder,
	UserCache,
	PreChecker,
	PostChecker,
	WebSecurityConfigurer,
	HTTPSecurityConfigurer,
	WebSecurityConfigurers,
	UserDetailsService,
	ClientDetailsService,
	DaoAuthenticationProvider,
)

// DaoAuthenticationProvider UsernamePasswordAuthenticationToken 认证提供者
var DaoAuthenticationProvider = wire.NewSet(wire.Struct(new(dao.AuthenticationProvider), "*"))

// Providers 所有认证提供者
func Providers(dao *dao.AuthenticationProvider, injector container.SecurityInjector) coreAuth.Providers {
	var providers coreAuth.Providers
	if len(injector.GetProviders()) != 0 {
		providers = injector.GetProviders()
	}
	providers = append(providers, dao)
	return providers
}

// PasswordEncoder encoder
func PasswordEncoder(injector container.SecurityInjector) password.Encoder {
	if injector.GetPasswordEncoder() != nil {
		return injector.GetPasswordEncoder()
	}
	return factory.CreateDelegatingPasswordEncoder()
}

// UserCache 用户缓存
func UserCache(injector container.SecurityInjector) userdetails.UserCache {
	if injector.GetUserCache() != nil {
		return injector.GetUserCache()
	}
	return cache.NewNilUserCache()
}

// PreChecker 前置检查器
func PreChecker(injector container.SecurityInjector) userdetails.PreChecker {
	if injector.GetPreChecker() != nil {
		return injector.GetPreChecker()
	}
	return dao.NewPreChecker()
}

// PostChecker 后置检查器
func PostChecker(injector container.SecurityInjector) userdetails.PostChecker {
	if injector.GetPostChecker() != nil {
		return injector.GetPostChecker()
	}
	return dao.NewPostChecker()
}

// WebSecurityConfigurer 默认配置
func WebSecurityConfigurer(injector container.SecurityInjector) security.WebSecurityConfigurer {
	return injector.GetWebSecurityConfigurer()
}

// HTTPSecurityConfigurer 默认配置
func HTTPSecurityConfigurer(injector container.SecurityInjector) security.HTTPSecurityConfigurer {
	return injector.GetHTTPSecurityConfigurer()
}

// WebSecurityConfigurers web 安全配置
func WebSecurityConfigurers(web security.WebSecurityConfigurer, http security.HTTPSecurityConfigurer, injector container.SecurityInjector) security.WebSecurityConfigurers {
	var configurers security.WebSecurityConfigurers
	if len(injector.GetWebSecurityConfigurers()) != 0 {
		configurers = injector.GetWebSecurityConfigurers()
	}
	configurers = append(configurers, config.NewWebSecurityConfigurerAdapter(web, http))
	return configurers
}

// UserDetailsService 用户详情服务
func UserDetailsService(injector container.SecurityInjector) userdetails.Service {
	if injector.GetUserDetailsService() != nil {
		return injector.GetUserDetailsService()
	}
	return NilUserDetailsService()
}

// ClientDetailsService 客户端详情服务
func ClientDetailsService(injector container.SecurityInjector) clientdetails.Service {
	if injector.GetClientDetailsService() != nil {
		return injector.GetClientDetailsService()
	}
	return NilClientDetails()
}
