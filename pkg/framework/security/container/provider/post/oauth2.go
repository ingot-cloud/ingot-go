package post

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/pre"
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
func TokenStore(sc container.SecurityContainerCombine) token.Store {
	return pre.TokenStore(sc.GetOAuth2Container().JwtAccessTokenConverter)
}

// JwtAccessTokenConverter 实例
func JwtAccessTokenConverter(sc container.SecurityContainerCombine) *store.JwtAccessTokenConverter {
	return pre.JwtAccessTokenConverter(sc.GetOAuth2Container().OAuth2Config, sc.GetOAuth2Container().AccessTokenConverter)
}

// AccessTokenConverter token转换器
func AccessTokenConverter(sc container.SecurityContainerCombine) token.AccessTokenConverter {
	return pre.AccessTokenConverter(sc.GetOAuth2Container().OAuth2Config, sc.GetOAuth2Container().UserAuthenticationConverter)
}

// UserAuthenticationConverter 默认实现
func UserAuthenticationConverter(sc container.SecurityContainerCombine) token.UserAuthenticationConverter {
	return pre.UserAuthenticationConverter()
}
