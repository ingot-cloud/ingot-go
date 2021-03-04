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

// CommonContainer 容器
var CommonContainer = wire.NewSet(wire.Struct(new(container.CommonContainer), "*"))

// CommonContainerFields 容器所有字段
var CommonContainerFields = wire.NewSet(
	PasswordEncoder,
	UserCache,
	PreChecker,
	PostChecker,
	UserDetailsService,
	ClientDetailsService,
	wire.Struct(new(WebSecurityConfigurersImpl)),
	wire.Bind(new(security.WebSecurityConfigurers), new(*WebSecurityConfigurersImpl)),
)

// WebSecurityConfigurersImpl 接口实现
type WebSecurityConfigurersImpl struct {
	configurers []security.WebSecurityConfigurer
	Injector    container.SecurityInjector
}

// Add 追加
func (web *WebSecurityConfigurersImpl) Add(c security.WebSecurityConfigurer) {
	if web.Injector.GetWebSecurityConfigurers() != nil {
		web.Injector.GetWebSecurityConfigurers().Add(c)
		return
	}
	web.configurers = append(web.configurers, c)
}

// Get 获取所有 WebSecurityConfigurer
func (web *WebSecurityConfigurersImpl) Get() []security.WebSecurityConfigurer {
	if web.Injector.GetWebSecurityConfigurers() != nil && len(web.Injector.GetWebSecurityConfigurers().Get()) != 0 {
		return web.Injector.GetWebSecurityConfigurers().Get()
	}
	return web.configurers
}

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
