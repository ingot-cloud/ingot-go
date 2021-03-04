package preset

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/model"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

// OAuth2Container OAuth2容器
var OAuth2Container = wire.NewSet(wire.Struct(new(container.OAuth2Container), "*"))

// OAuth2ContainerFields OAuth2容器所有字段
var OAuth2ContainerFields = wire.NewSet(
	DefaultTokenServices,
	TokenStore,
	JwtAccessTokenConverter,
	AccessTokenConverter,
	UserAuthenticationConverter,
)

// DefaultTokenServices 默认的服务
func DefaultTokenServices(config model.OAuth2, tokenStore token.Store) *token.DefaultTokenServices {
	service := token.NewDefaultTokenServices(tokenStore)
	service.ReuseRefreshToken = config.AuthorizationServer.ReuseRefreshToken
	service.SupportRefreshToken = config.AuthorizationServer.SupportRefreshToken
	return service
}

// TokenStore 实例
func TokenStore(converter *store.JwtAccessTokenConverter) token.Store {
	return store.NewJwtTokenStore(converter)
}

// JwtAccessTokenConverter 实例
func JwtAccessTokenConverter(config model.OAuth2, tokenConverter token.AccessTokenConverter) *store.JwtAccessTokenConverter {
	var method jwt.SigningMethod
	switch config.Jwt.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}
	signingKey := []byte(config.Jwt.SigningKey)
	keyfunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.InvalidToken("Token invalid")
		}
		return []byte(config.Jwt.SigningKey), nil
	}

	return store.NewJwtAccessTokenConverter(tokenConverter, method, signingKey, keyfunc)
}

// AccessTokenConverter token转换器
func AccessTokenConverter(config model.OAuth2, userConverter token.UserAuthenticationConverter) token.AccessTokenConverter {
	converter := token.NewDefaultAccessTokenConverter()
	converter.IncludeGrantType = config.IncludeGrantType
	converter.UserAuthenticationConverter = userConverter
	return converter
}

// UserAuthenticationConverter 默认实现
func UserAuthenticationConverter() token.UserAuthenticationConverter {
	return token.NewDefaultUserAuthenticationConverter()
}
