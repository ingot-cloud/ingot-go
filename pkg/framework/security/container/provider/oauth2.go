package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/preset"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

// OAuth2Container OAuth2容器
var OAuth2Container = wire.NewSet(wire.Struct(new(container.OAuth2Container), "*"))

// OAuth2ContainerFields OAuth2容器所有字段
var OAuth2ContainerFields = wire.NewSet(
	TokenStore,
	JwtAccessTokenConverter,
	AccessTokenConverter,
	UserAuthenticationConverter,
)

// TokenStore 实例
func TokenStore(converter *store.JwtAccessTokenConverter, injector container.SecurityInjector) token.Store {
	if injector.GetTokenStore() != nil {
		return injector.GetTokenStore()
	}
	return preset.TokenStore(converter)
}

// JwtAccessTokenConverter 实例
func JwtAccessTokenConverter(config config.OAuth2, tokenConverter token.AccessTokenConverter) *store.JwtAccessTokenConverter {
	return preset.JwtAccessTokenConverter(config, tokenConverter)
}

// AccessTokenConverter token转换器
func AccessTokenConverter(config config.OAuth2, userConverter token.UserAuthenticationConverter, injector container.SecurityInjector) token.AccessTokenConverter {
	if injector.GetAccessTokenConverter() != nil {
		return injector.GetAccessTokenConverter()
	}
	return preset.AccessTokenConverter(config, userConverter)
}

// UserAuthenticationConverter 默认实现
func UserAuthenticationConverter(injector container.SecurityInjector) token.UserAuthenticationConverter {
	if injector.GetUserAuthenticationConverter() != nil {
		return injector.GetUserAuthenticationConverter()
	}
	return preset.UserAuthenticationConverter()
}
