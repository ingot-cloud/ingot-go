package store

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// todo

// JwtTokenStore TokenStore jwt 实现
type JwtTokenStore struct {
	JwtTokenEnhancer *JwtAccessTokenConverter
}

// NewJwtTokenStore 创建 JwtTokenStore
func NewJwtTokenStore(converter *JwtAccessTokenConverter) *JwtTokenStore {
	return &JwtTokenStore{
		JwtTokenEnhancer: converter,
	}
}

// ReadAuthentication 根据token读取身份验证信息
func (store *JwtTokenStore) ReadAuthentication(accessToken token.OAuth2AccessToken) (*authentication.OAuth2Authentication, error) {
	return store.ReadAuthenticationWith(accessToken.GetValue())
}

// ReadAuthenticationWith 根据token读取身份验证信息
func (store *JwtTokenStore) ReadAuthenticationWith(tokenValue string) (*authentication.OAuth2Authentication, error) {
	info, err := store.JwtTokenEnhancer.Decode(tokenValue)
	if err != nil {
		return nil, err
	}
	return store.JwtTokenEnhancer.ExtractAuthentication(info)
}

// StoreAccessToken 存储访问令牌
func (store *JwtTokenStore) StoreAccessToken(accessToken token.OAuth2AccessToken, authentication *authentication.OAuth2Authentication) {

}

// ReadAccessToken 读取访问令牌
func (store *JwtTokenStore) ReadAccessToken(tokenValue string) (token.OAuth2AccessToken, error) {
	accessToken, err := store.convertAccessToken(tokenValue)
	if err != nil {
		return nil, err
	}
	if store.JwtTokenEnhancer.IsRefreshToken(accessToken) {
		return nil, errors.InvalidToken("Encoded token is a refresh token")
	}

	return accessToken, nil
}

// RemoveAccessToken 移除访问令牌
func (store *JwtTokenStore) RemoveAccessToken(accessToken token.OAuth2AccessToken) {

}

// StoreRefreshToken 存储刷新令牌
func (store *JwtTokenStore) StoreRefreshToken(accessToken token.OAuth2RefreshToken, authentication *authentication.OAuth2Authentication) {

}

// ReadRefreshToken 读取刷新令牌
func (store *JwtTokenStore) ReadRefreshToken(tokenValue string) (token.OAuth2RefreshToken, error) {
	encodedRefreshToken, err := store.convertAccessToken(tokenValue)
	if err != nil {
		return nil, err
	}

	return store.createRefreshToken(encodedRefreshToken)
}

// ReadAuthenticationForRefreshToken 通过刷新令牌读取身份验证信息
func (store *JwtTokenStore) ReadAuthenticationForRefreshToken(refreshToken token.OAuth2RefreshToken) (*authentication.OAuth2Authentication, error) {
	return store.ReadAuthenticationWith(refreshToken.GetRefreshTokenValue())
}

// RemoveRefreshToken 移除刷新令牌
func (store *JwtTokenStore) RemoveRefreshToken(refreshToken token.OAuth2RefreshToken) {

}

// RemoveAccessTokenUsingRefreshToken 通过刷新令牌移除访问令牌
func (store *JwtTokenStore) RemoveAccessTokenUsingRefreshToken(refreshToken token.OAuth2RefreshToken) {

}

// GetAccessToken 通过身份验证信息获取访问令牌
func (store *JwtTokenStore) GetAccessToken(authentication *authentication.OAuth2Authentication) (token.OAuth2AccessToken, error) {
	return nil, nil
}

// FindTokensByClientIDAndUserName 通过clientID和用户名获取所有访问令牌
func (store *JwtTokenStore) FindTokensByClientIDAndUserName(clientID string, username string) ([]token.OAuth2AccessToken, error) {
	return nil, nil
}

// FindTokensByClientID 通过clientID获取所有访问令牌
func (store *JwtTokenStore) FindTokensByClientID(clientID string) ([]token.OAuth2AccessToken, error) {
	return nil, nil
}

func (store *JwtTokenStore) convertAccessToken(tokenValue string) (token.OAuth2AccessToken, error) {
	info, err := store.JwtTokenEnhancer.Decode(tokenValue)
	if err != nil {
		return nil, err
	}
	return store.JwtTokenEnhancer.ExtractAccessToken(tokenValue, info)
}

func (store *JwtTokenStore) createRefreshToken(encodedRefreshToken token.OAuth2AccessToken) (token.OAuth2RefreshToken, error) {
	if !store.JwtTokenEnhancer.IsRefreshToken(encodedRefreshToken) {
		return nil, errors.InvalidToken("Encoded token is not a refresh token")
	}
	if !encodedRefreshToken.GetExpiration().IsZero() {
		refreshToken := token.NewDefaultExpiringOAuth2RefreshToken(encodedRefreshToken.GetValue(), encodedRefreshToken.GetExpiration())
		return refreshToken, nil
	}

	return token.NewDefaultOAuth2RefreshToken(encodedRefreshToken.GetValue()), nil
}
