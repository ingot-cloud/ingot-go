package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
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
	wire.Struct(new(WebSecurityConfigurersImpl), "*"),
	wire.Bind(new(security.WebSecurityConfigurers), new(*WebSecurityConfigurersImpl)),
)

// WebSecurityConfigurersImpl 接口实现
type WebSecurityConfigurersImpl struct {
	SC container.SecurityContainerCombine
}

// Add 追加
func (web *WebSecurityConfigurersImpl) Add(c security.WebSecurityConfigurer) {
	web.SC.GetCommonContainer().WebSecurityConfigurers.Add(c)
}

// Get 获取所有 WebSecurityConfigurer
func (web *WebSecurityConfigurersImpl) Get() []security.WebSecurityConfigurer {
	return web.SC.GetCommonContainer().WebSecurityConfigurers.Get()
}

// PasswordEncoder encoder
func PasswordEncoder(sc container.SecurityContainerCombine) password.Encoder {
	return sc.GetCommonContainer().PasswordEncoder
}

// UserCache 用户缓存
func UserCache(sc container.SecurityContainerCombine) userdetails.UserCache {
	return sc.GetCommonContainer().UserCache
}

// PreChecker 前置检查器
func PreChecker(sc container.SecurityContainerCombine) userdetails.PreChecker {
	return sc.GetCommonContainer().PreChecker
}

// PostChecker 后置检查器
func PostChecker(sc container.SecurityContainerCombine) userdetails.PostChecker {
	return sc.GetCommonContainer().PostChecker
}

// UserDetailsService 用户详情服务
func UserDetailsService(sc container.SecurityContainerCombine) userdetails.Service {
	return sc.GetCommonContainer().UserDetailsService
}

// ClientDetailsService 客户端详情服务
func ClientDetailsService(sc container.SecurityContainerCombine) clientdetails.Service {
	return sc.GetCommonContainer().ClientDetailsService
}
