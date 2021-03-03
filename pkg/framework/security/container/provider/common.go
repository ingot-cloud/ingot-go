package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/preset"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// Common 容器
var Common = wire.NewSet(wire.Struct(new(container.Common), "*"))

// CommonFields 容器所有字段
var CommonFields = wire.NewSet(
	PasswordEncoder,
	UserCache,
	PreChecker,
	PostChecker,
	WebSecurityConfigurer,
	HTTPSecurityConfigurer,
	WebSecurityConfigurers,
	UserDetailsService,
	ClientDetailsService,
)

// PasswordEncoder encoder
func PasswordEncoder(injector container.SecurityInjector) password.Encoder {
	if injector.GetPasswordEncoder() != nil {
		return injector.GetPasswordEncoder()
	}
	return preset.PasswordEncoder()
}

// UserCache 用户缓存
func UserCache(injector container.SecurityInjector) userdetails.UserCache {
	if injector.GetUserCache() != nil {
		return injector.GetUserCache()
	}
	return preset.UserCache()
}

// PreChecker 前置检查器
func PreChecker(injector container.SecurityInjector) userdetails.PreChecker {
	if injector.GetPreChecker() != nil {
		return injector.GetPreChecker()
	}
	return preset.PreChecker()
}

// PostChecker 后置检查器
func PostChecker(injector container.SecurityInjector) userdetails.PostChecker {
	if injector.GetPostChecker() != nil {
		return injector.GetPostChecker()
	}
	return preset.PostChecker()
}

// WebSecurityConfigurer 默认配置
func WebSecurityConfigurer(injector container.SecurityInjector) security.WebSecurityConfigurer {
	if injector.GetWebSecurityConfigurer() != nil {
		return injector.GetWebSecurityConfigurer()
	}
	return preset.WebSecurityConfigurer()
}

// HTTPSecurityConfigurer 默认配置
func HTTPSecurityConfigurer(injector container.SecurityInjector) security.HTTPSecurityConfigurer {
	if injector.GetHTTPSecurityConfigurer() != nil {
		return injector.GetHTTPSecurityConfigurer()
	}
	return preset.HTTPSecurityConfigurer()
}

// WebSecurityConfigurers web 安全配置
func WebSecurityConfigurers(web security.WebSecurityConfigurer, http security.HTTPSecurityConfigurer, injector container.SecurityInjector) security.WebSecurityConfigurers {
	if len(injector.GetWebSecurityConfigurers()) != 0 {
		return injector.GetWebSecurityConfigurers()
	}
	return preset.WebSecurityConfigurers(web, http)
}

// UserDetailsService 用户详情服务
func UserDetailsService(injector container.SecurityInjector) userdetails.Service {
	if injector.GetUserDetailsService() != nil {
		return injector.GetUserDetailsService()
	}
	return preset.UserDetailsService()
}

// ClientDetailsService 客户端详情服务
func ClientDetailsService(injector container.SecurityInjector) clientdetails.Service {
	if injector.GetClientDetailsService() != nil {
		return injector.GetClientDetailsService()
	}
	return preset.ClientDetailsService()
}
