package provider

import (
	"github.com/google/wire"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails/cache"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/factory"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
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

// DaoAuthenticationProviderSet UsernamePasswordAuthenticationToken 认证提供者
var DaoAuthenticationProviderSet = wire.NewSet(wire.Struct(new(dao.AuthenticationProvider), "*"))
