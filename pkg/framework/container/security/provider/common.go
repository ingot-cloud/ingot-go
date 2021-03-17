package provider

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/container/security/provider/null"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails/cache"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/factory"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
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

// UserDetailsService 用户详情服务
func UserDetailsService() userdetails.Service {
	return null.UserDetailsService()
}

// ClientDetailsService 客户端详情服务
func ClientDetailsService() clientdetails.Service {
	return null.ClientDetails()
}
