package provider

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

// TokenStore 实例
func TokenStore(converter *store.JwtAccessTokenConverter) token.Store {
	return store.NewJwtTokenStore(converter)
}

// JwtAccessTokenConverter 实例
func JwtAccessTokenConverter(config config.OAuth2, tokenConverter token.AccessTokenConverter) *store.JwtAccessTokenConverter {
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
func AccessTokenConverter(config config.OAuth2, userConverter token.UserAuthenticationConverter) token.AccessTokenConverter {
	converter := token.NewDefaultAccessTokenConverter(userConverter)
	converter.IncludeGrantType = config.IncludeGrantType
	return converter
}

// UserAuthenticationConverter 默认实现
func UserAuthenticationConverter() token.UserAuthenticationConverter {
	return token.NewDefaultUserAuthenticationConverter()
}
