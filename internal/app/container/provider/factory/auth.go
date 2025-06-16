package factory

// NewAuthentication for auth
// func NewAuthentication(config *config.Config) (security.Authentication, func(), error) {
// authCfg := config.Auth.Jwt
// redisCfg := config.Redis

// store := jwtAuth.NewTokenStore(&store.RedisParams{
// 	Address:   redisCfg.Address,
// 	DB:        redisCfg.DB,
// 	Password:  redisCfg.Password,
// 	KeyPrefix: redisCfg.KeyPrefix + "AUTH:",
// 	SSL:       redisCfg.SSL,
// })

// var method jwt.SigningMethod
// switch authCfg.SigningMethod {
// case "HS256":
// 	method = jwt.SigningMethodHS256
// case "HS384":
// 	method = jwt.SigningMethodHS384
// default:
// 	method = jwt.SigningMethodHS512
// }

// authParms := &jwtAuth.Params{
// 	SigningMethod: method,
// 	SigningKey:    []byte(authCfg.SigningKey),
// 	Keyfunc: func(t *jwt.Token) (any, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, security.ErrInvalidToken
// 		}
// 		return []byte(authCfg.SigningKey), nil
// 	},
// 	// Expired:   authCfg.Expired,
// 	TokenType: "Bearer",
// }

// auth := jwtAuth.NewAuthentication(store, authParms)
// cleanFunc := func() {
// 	auth.Release()
// }
// 	return nil, nil, nil
// }
