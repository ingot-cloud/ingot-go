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
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

// SecurityContainer 安全容器
var SecurityContainer = wire.NewSet(wire.Struct(new(container.SecurityContainer), "*"))

// SecurityContainerFields 安全容器所有字段
var SecurityContainerFields = wire.NewSet(
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

// WebSecurityConfigurers web 安全配置
func WebSecurityConfigurers(web security.WebSecurityConfigurer, http security.HTTPSecurityConfigurer) security.WebSecurityConfigurers {
	var configurers security.WebSecurityConfigurers
	configurers = append(configurers, config.NewWebSecurityConfigurerAdapter(web, http))
	return configurers
}

// UserDetailsService 用户详情服务
func UserDetailsService() userdetails.Service {
	return NilUserDetailsService()
}

// ClientDetailsService 客户端详情服务
func ClientDetailsService() clientdetails.Service {
	return NilClientDetails()
}
