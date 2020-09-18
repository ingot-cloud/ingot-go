package provider

import (
	"ingot/config"
	"ingot/core/security"
	jwtAuth "ingot/core/security/jwt"
	"ingot/core/store"

	jwt "github.com/dgrijalva/jwt-go"
)

// BuildAuthentication for auth
func BuildAuthentication() (security.Authentication, func(), error) {
	authCfg := config.CONFIG.Auth.Jwt
	redisCfg := config.CONFIG.Redis

	store := jwtAuth.NewTokenStore(&store.RedisParams{
		Address:   redisCfg.Address,
		DB:        3,
		Password:  redisCfg.Password,
		KeyPrefix: "AUTH:",
	})

	var method jwt.SigningMethod
	switch authCfg.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}

	authParms := &jwtAuth.Params{
		SigningMethod: method,
		SigningKey:    []byte(authCfg.SigningKey),
		Keyfunc: func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, security.ErrInvalidToken
			}
			return []byte(authCfg.SigningKey), nil
		},
		Expired:   authCfg.Expired,
		TokenType: "Bearer",
	}

	auth := jwtAuth.NewAuthentication(store, authParms)
	cleanFunc := func() {
		auth.Release()
	}
	return auth, cleanFunc, nil
}
