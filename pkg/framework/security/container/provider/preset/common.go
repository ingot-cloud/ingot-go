package preset

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails/cache"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/factory"
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
	WebSecurityConfigurer,
	HTTPSecurityConfigurer,
	UserDetailsService,
	ClientDetailsService,
	wire.Struct(new(WebSecurityConfigurersImpl)),
	wire.Bind(new(security.WebSecurityConfigurers), new(*WebSecurityConfigurersImpl)),
)

// WebSecurityConfigurersImpl 接口实现
type WebSecurityConfigurersImpl struct {
	configurers []security.WebSecurityConfigurer
}

// Add 追加
func (web *WebSecurityConfigurersImpl) Add(c security.WebSecurityConfigurer) {
	web.configurers = append(web.configurers, c)
}

// Get 获取所有 WebSecurityConfigurer
func (web *WebSecurityConfigurersImpl) Get() []security.WebSecurityConfigurer {
	return web.configurers
}

// PasswordEncoder encoder
func PasswordEncoder() password.Encoder {
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

// WebSecurityConfigurer 默认配置
func WebSecurityConfigurer() security.WebSecurityConfigurer {
	return nil
}

// HTTPSecurityConfigurer 默认配置
func HTTPSecurityConfigurer() security.HTTPSecurityConfigurer {
	return nil
}

// UserDetailsService 用户详情服务
func UserDetailsService() userdetails.Service {
	return NilUserDetailsService()
}

// ClientDetailsService 客户端详情服务
func ClientDetailsService() clientdetails.Service {
	return NilClientDetails()
}
