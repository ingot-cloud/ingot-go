package jwt

import (
	"context"
	"ingot/core/security"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// NewAuthentication for jwt
func NewAuthentication(store security.TokenStore, params *Params) security.Authentication {
	return &Authentication{
		store:  store,
		params: params,
	}
}

// Claims custom
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Params for create
type Params struct {
	SigningMethod jwt.SigningMethod
	SigningKey    interface{}
	Keyfunc       jwt.Keyfunc
	Expired       int
	TokenType     string
}

// Authentication impl security.Authentication
type Authentication struct {
	params *Params
	store  security.TokenStore
}

// GenerateToken 生成令牌
func (a *Authentication) GenerateToken(ctx context.Context, user security.User) (security.AccessToken, error) {
	now := time.Now()
	expiration := now.Add(time.Duration(a.params.Expired) * time.Second).Unix()
	token := jwt.NewWithClaims(a.params.SigningMethod, &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: expiration,
			NotBefore: now.Unix(),
			Subject:   user.ID,
		},
	})

	accessToken, err := token.SignedString(a.params.SigningKey)
	if err != nil {
		return nil, err
	}

	return &AccessToken{
		AccessToken: accessToken,
		Expiration:  expiration,
		TokenType:   a.params.TokenType,
	}, nil
}

// ParseUser 解析用户
func (a *Authentication) ParseUser(ctx context.Context, accessToken string) (*security.User, error) {
	if accessToken == "" {
		return nil, security.ErrInvalidToken
	}

	claims, err := a.parse(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 0 {
		return nil, security.ErrExpiredToken
	}

	err = a.ensureStore(func(store security.TokenStore) error {
		if exists, err := store.Check(ctx, accessToken); err != nil {
			return err
		} else if !exists {
			return security.ErrInvalidToken
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &security.User{
		ID:       claims.Subject,
		Username: claims.Username,
	}, nil
}

// DestroyToken 销毁令牌
func (a *Authentication) DestroyToken(ctx context.Context, accessToken string) error {
	if accessToken == "" {
		return nil
	}

	_, err := a.parse(ctx, accessToken)
	if err != nil {
		return err
	}

	return a.ensureStore(func(store security.TokenStore) error {
		store.Remove(ctx, accessToken)
		return nil
	})
}

// Release 释放资源
func (a *Authentication) Release() error {
	return a.ensureStore(func(store security.TokenStore) error {
		return store.Close()
	})
}

func (a *Authentication) parse(ctx context.Context, accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, a.params.Keyfunc)
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, security.ErrInvalidToken
	}

	return token.Claims.(*Claims), nil
}

func (a *Authentication) ensureStore(fn func(security.TokenStore) error) error {
	if store := a.store; store != nil {
		return fn(store)
	}
	return nil
}
