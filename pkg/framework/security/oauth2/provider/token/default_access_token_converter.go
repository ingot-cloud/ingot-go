package token

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/authority"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/request"
)

// DefaultAccessTokenConverter 默认实现
type DefaultAccessTokenConverter struct {
	UserAuthenticationConverter *DefaultUserAuthenticationConverter
	IncludeGrantType            bool
}

// NewDefaultAccessTokenConverter 实例化
func NewDefaultAccessTokenConverter() *DefaultAccessTokenConverter {
	return &DefaultAccessTokenConverter{
		UserAuthenticationConverter: NewDefaultUserAuthenticationConverter(),
		IncludeGrantType:            false,
	}
}

// ConvertAccessToken 返回访问令牌映射内容
func (converter *DefaultAccessTokenConverter) ConvertAccessToken(token OAuth2AccessToken, authentication *authentication.OAuth2Authentication) (map[string]interface{}, error) {
	response := make(map[string]interface{})
	clientToken := authentication.GetOAuth2Request()

	if !authentication.IsClientOnly() {
		r, err := converter.getUserAuthenticationConverter().ConvertUserAuthentication(authentication.UserAuthentication)
		if err != nil {
			return nil, err
		}
		for k, v := range r {
			response[k] = v
		}
	} else {
		authorities := clientToken.GetAuthorities()
		if len(authorities) != 0 {
			response[string(constants.TokenAuthorities)] = authorities
		}
	}

	// scope
	response[string(constants.TokenScope)] = token.GetScope()

	// jti
	if jti, ok := token.GetAdditionalInformation()[string(constants.TokenJti)]; ok {
		response[string(constants.TokenJti)] = jti
	}

	// exp
	if exp := token.GetExpiration(); !utils.TimeIsNil(exp) {
		// 单位秒
		response[string(constants.TokenExp)] = exp.Unix()
	}

	// additional
	for k, v := range token.GetAdditionalInformation() {
		response[k] = v
	}

	// grant_type
	if converter.IncludeGrantType && clientToken.GetGrantType() != "" {
		response[string(constants.TokenGrantType)] = clientToken.GetGrantType()
	}

	// client_id
	response[string(constants.TokenClientID)] = clientToken.GetClientID()

	// aud
	if rids := clientToken.GetResourceIDs(); len(rids) != 0 {
		response[string(constants.TokenAud)] = rids
	}

	return response, nil
}

// ExtractAccessToken 根据token value和映射内容提取访问令牌
func (converter *DefaultAccessTokenConverter) ExtractAccessToken(token string, mapInfo map[string]interface{}) (OAuth2AccessToken, error) {
	accessToken := NewDefaultOAuth2AccessToken(token)
	info := make(map[string]interface{})
	for k, v := range mapInfo {
		info[k] = v
	}
	delete(info, string(constants.TokenExp))
	delete(info, string(constants.TokenAud))
	delete(info, string(constants.TokenClientID))
	delete(info, string(constants.TokenScope))
	if jti, ok := mapInfo[string(constants.TokenJti)]; ok {
		info[string(constants.TokenJti)] = jti
	}
	if exp, ok := mapInfo[string(constants.TokenExp)]; ok {
		if val, ok := exp.(int64); ok {
			accessToken.Expiration = time.Unix(val, 0)
		}
	}
	accessToken.Scope = converter.extractScope(mapInfo)
	accessToken.AdditionalInformation = info
	return accessToken, nil
}

// ExtractAuthentication 根据token映射信息提取身份验证信息
func (converter *DefaultAccessTokenConverter) ExtractAuthentication(mapInfo map[string]interface{}) (*authentication.OAuth2Authentication, error) {
	parameters := make(map[string]string)
	scope := converter.extractScope(mapInfo)
	user, err := converter.getUserAuthenticationConverter().ExtractAuthentication(mapInfo)
	if err != nil {
		return nil, err
	}
	clientID, ok := mapInfo[string(constants.TokenClientID)].(string)
	if ok {
		parameters[string(constants.TokenClientID)] = clientID
	}
	if converter.IncludeGrantType {
		if grantType, ok := mapInfo[string(constants.TokenGrantType)].(string); ok {
			parameters[string(constants.TokenGrantType)] = grantType
		}
	}

	resourceIDs := converter.getAudience(mapInfo)
	authorities := authority.CreateAuthorityList(mapInfo[string(constants.TokenAuthorities)])

	request := request.NewOAuth2Request(parameters, clientID, scope)
	request.ResourceIDs = resourceIDs
	request.Authorities = authorities
	request.Approved = true

	return &authentication.OAuth2Authentication{
		StoredRequest:      request,
		UserAuthentication: user,
	}, nil
}

func (converter *DefaultAccessTokenConverter) getUserAuthenticationConverter() *DefaultUserAuthenticationConverter {
	if converter.UserAuthenticationConverter == nil {
		converter.UserAuthenticationConverter = NewDefaultUserAuthenticationConverter()
	}
	return converter.UserAuthenticationConverter
}

func (converter *DefaultAccessTokenConverter) extractScope(mapInfo map[string]interface{}) []string {
	if scope, ok := mapInfo[string(constants.TokenScope)]; ok {
		switch value := scope.(type) {
		case string:
			return []string{value}
		case []string:
			return value
		}
	}
	return nil
}

func (converter *DefaultAccessTokenConverter) getAudience(mapInfo map[string]interface{}) []string {
	auds, ok := mapInfo[string(constants.TokenAud)]
	if !ok {
		return nil
	}
	switch value := auds.(type) {
	case string:
		return []string{value}
	case []string:
		return value
	default:
		return nil
	}
}
