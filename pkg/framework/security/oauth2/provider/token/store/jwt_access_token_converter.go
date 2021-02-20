package store

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils/maputil"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// JwtAccessTokenConverter jwt和oauth2身份验证信息转换器
type JwtAccessTokenConverter struct {
	tokenConverter token.AccessTokenConverter
	SigningMethod  jwt.SigningMethod
	SigningKey     interface{}
	Keyfunc        jwt.Keyfunc
}

// NewJwtAccessTokenConverter 实例化
func NewJwtAccessTokenConverter(method jwt.SigningMethod, signingKey interface{}) *JwtAccessTokenConverter {
	return &JwtAccessTokenConverter{
		tokenConverter: token.NewDefaultAccessTokenConverter(),
		SigningMethod:  method,
		SigningKey:     signingKey,
	}
}

// ConvertAccessToken 返回访问令牌映射内容
func (c *JwtAccessTokenConverter) ConvertAccessToken(accessToken token.OAuth2AccessToken, authentication *authentication.OAuth2Authentication) (map[string]interface{}, error) {
	return c.tokenConverter.ConvertAccessToken(accessToken, authentication)
}

// ExtractAccessToken 根据token value和映射内容提取访问令牌
func (c *JwtAccessTokenConverter) ExtractAccessToken(accessToken string, mapInfo map[string]interface{}) (token.OAuth2AccessToken, error) {
	return c.tokenConverter.ExtractAccessToken(accessToken, mapInfo)
}

// ExtractAuthentication 根据token映射信息提取身份验证信息
func (c *JwtAccessTokenConverter) ExtractAuthentication(mapInfo map[string]interface{}) (*authentication.OAuth2Authentication, error) {
	return c.tokenConverter.ExtractAuthentication(mapInfo)
}

// Enhance Enhancer 接口实现
func (c *JwtAccessTokenConverter) Enhance(accessToken token.OAuth2AccessToken, authentication *authentication.OAuth2Authentication) (token.OAuth2AccessToken, error) {
	result := token.NewDefaultOAuth2AccessTokenWith(accessToken)
	info := maputil.CopyStringInterfaceMap(accessToken.GetAdditionalInformation())
	tokenID, ok := info[string(constants.TokenJti)]
	if !ok {
		tokenID = accessToken.GetValue()
		info[string(constants.TokenJti)] = tokenID
	}

	result.AdditionalInformation = info
	encode, err := c.Encode(result, authentication)
	if err != nil {
		return nil, err
	}
	result.Value = encode

	refreshToken := result.GetRefreshToken()
	if refreshToken != nil {
		encodedRefreshToken := token.NewDefaultOAuth2AccessTokenWith(accessToken)
		encodedRefreshToken.Value = refreshToken.GetRefreshTokenValue()
		encodedRefreshToken.Expiration = time.Time{}

		// 判断 token value 中是否可以解析出 jti
		// 如果获取到的 refreshToken 是 jwt 那么可以解析出来 jti，解析后需要替换原值
		if t, ok := c.getTokenID(refreshToken.GetRefreshTokenValue()); ok {
			encodedRefreshToken.Value = t
		}
		refreshTokenInfo := maputil.CopyStringInterfaceMap(accessToken.GetAdditionalInformation())
		refreshTokenInfo[string(constants.TokenJti)] = encodedRefreshToken.GetValue()
		refreshTokenInfo[string(constants.TokenAti)] = tokenID
		encodedRefreshToken.AdditionalInformation = refreshTokenInfo

		refreshEncode, err := c.Encode(encodedRefreshToken, authentication)
		if err != nil {
			return nil, err
		}
		var newRefreshToken token.OAuth2RefreshToken
		newRefreshToken = token.NewDefaultOAuth2RefreshToken(refreshEncode)
		if expiring, ok := refreshToken.(token.ExpiringOAuth2RefreshToken); ok {
			encodedRefreshToken.Expiration = expiring.GetExpiration()
			refreshEncode, err = c.Encode(encodedRefreshToken, authentication)
			if err != nil {
				return nil, err
			}
			newRefreshToken = token.NewDefaultExpiringOAuth2RefreshToken(refreshEncode, expiring.GetExpiration())
		}
		result.RefreshToken = newRefreshToken
	}

	return result, nil
}

// SetAccessTokenConverter 设置访问令牌转换器
func (c *JwtAccessTokenConverter) SetAccessTokenConverter(tokenConverter token.AccessTokenConverter) {
	c.tokenConverter = tokenConverter
}

// GetAccessTokenConverter 获取访问令牌转换器
func (c *JwtAccessTokenConverter) GetAccessTokenConverter() token.AccessTokenConverter {
	if c.tokenConverter == nil {
		c.tokenConverter = &token.DefaultAccessTokenConverter{}
	}
	return c.tokenConverter
}

// Encode 编码
func (c *JwtAccessTokenConverter) Encode(accessToken token.OAuth2AccessToken, auth *authentication.OAuth2Authentication) (string, error) {
	tokenInfo, err := c.GetAccessTokenConverter().ConvertAccessToken(accessToken, auth)
	if err != nil {
		return "", err
	}

	// 设置自定义 Claims
	mapClaims := jwt.MapClaims{}
	for k, v := range tokenInfo {
		mapClaims[k] = v
	}

	jwtToken := jwt.NewWithClaims(c.SigningMethod, mapClaims)

	tokenValue, err := jwtToken.SignedString(c.SigningKey)
	if err != nil {
		return "", nil
	}

	return tokenValue, nil
}

// Decode 解码
func (c *JwtAccessTokenConverter) Decode(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, c.Keyfunc)
	if err != nil {
		return nil, errors.InvalidToken(err.Error())
	} else if !token.Valid {
		return nil, errors.ErrInvalidToken
	}

	return token.Claims.(jwt.MapClaims), nil
}

// IsRefreshToken 判断是否为 RefreshToken，如果包含 ati 那么为 RefreshToken
func (c *JwtAccessTokenConverter) IsRefreshToken(token token.OAuth2AccessToken) bool {
	info := token.GetAdditionalInformation()
	_, ok := info[string(constants.TokenAti)]
	return ok
}

func (c *JwtAccessTokenConverter) getTokenID(tokenString string) (string, bool) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, c.Keyfunc)
	if err != nil || !token.Valid {
		return "", false
	}

	tokenID, ok := token.Claims.(jwt.MapClaims)[string(constants.TokenJti)]
	return tokenID.(string), ok
}
