package store

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// todo

// JwtTokenStore TokenStore jwt 实现
type JwtTokenStore struct {
	JwtAccessTokenConverter *JwtAccessTokenConverter
}

// ReadAuthentication 根据token读取身份验证信息
func (store *JwtTokenStore) ReadAuthentication(accessToken token.OAuth2AccessToken) (*authentication.OAuth2Authentication, error) {
	return nil, nil
}

// ReadAuthenticationWith 根据token读取身份验证信息
func (store *JwtTokenStore) ReadAuthenticationWith(tokenValue string) (*authentication.OAuth2Authentication, error) {
	return nil, nil
}

// StoreAccessToken 存储访问令牌
func (store *JwtTokenStore) StoreAccessToken(accessToken token.OAuth2AccessToken, authentication *authentication.OAuth2Authentication) {

}

// ReadAccessToken 读取访问令牌
func (store *JwtTokenStore) ReadAccessToken(tokenValue string) (token.OAuth2AccessToken, error) {
	return nil, nil
}

// RemoveAccessToken 移除访问令牌
func (store *JwtTokenStore) RemoveAccessToken(accessToken token.OAuth2AccessToken) {

}

// StoreRefreshToken 存储刷新令牌
func (store *JwtTokenStore) StoreRefreshToken(accessToken token.OAuth2RefreshToken, authentication *authentication.OAuth2Authentication) {

}

// ReadRefreshToken 读取刷新令牌
func (store *JwtTokenStore) ReadRefreshToken(tokenValue string) (token.OAuth2RefreshToken, error) {
	return nil, nil
}

// ReadAuthenticationForRefreshToken 通过刷新令牌读取身份验证信息
func (store *JwtTokenStore) ReadAuthenticationForRefreshToken(refreshToken token.OAuth2RefreshToken) (*authentication.OAuth2Authentication, error) {
	return nil, nil
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
