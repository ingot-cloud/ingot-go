package token

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/authority"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
)

// DefaultUserAuthenticationConverter 用户身份信息转换器默认实现
type DefaultUserAuthenticationConverter struct {
	DefaultAuthorities []core.GrantedAuthority
	UserDetailsService userdetails.Service
}

// NewDefaultUserAuthenticationConverter 实例DefaultUserAuthenticationConverter
func NewDefaultUserAuthenticationConverter() *DefaultUserAuthenticationConverter {
	return &DefaultUserAuthenticationConverter{}
}

// SetUserDetailsService 设置 UserDetailsService
func (converter *DefaultUserAuthenticationConverter) SetUserDetailsService(service userdetails.Service) {
	converter.UserDetailsService = service
}

// SetDefaultAuthorities 设置默认权限
func (converter *DefaultUserAuthenticationConverter) SetDefaultAuthorities(defaultAuthorities []string) {
	converter.DefaultAuthorities = authority.CreateAuthorityList(defaultAuthorities)
}

// ConvertUserAuthentication 在身份验证信息中提取访问令牌使用的信息
func (converter *DefaultUserAuthenticationConverter) ConvertUserAuthentication(auth core.Authentication) (map[string]interface{}, error) {
	if auth == nil {
		return nil, nil
	}
	response := make(map[string]interface{})
	response[string(constants.TokenUsername)] = auth.GetName(auth)
	authorities := auth.GetAuthorities()
	if len(authorities) != 0 {
		response[string(constants.TokenAuthorities)] = authorities
	}
	return response, nil
}

// ExtractAuthentication 从map中提取身份验证信息
func (converter *DefaultUserAuthenticationConverter) ExtractAuthentication(mapInfo map[string]interface{}) (core.Authentication, error) {
	principal, ok := mapInfo[string(constants.TokenUsername)]
	if ok {
		authorities := converter.getAuthorities(mapInfo)
		if converter.UserDetailsService != nil {
			username, ok := principal.(string)
			if ok {
				user, err := converter.UserDetailsService.LoadUserByUsername(username)
				if err != nil {
					return nil, err
				}
				authorities = user.GetAuthorities()
				principal = user
			}
		}
		return authentication.NewAuthenticatedUsernamePasswordAuthToken(principal, "N/A", authorities), nil
	}
	return nil, nil
}

func (converter *DefaultUserAuthenticationConverter) getAuthorities(mapInfo map[string]interface{}) []core.GrantedAuthority {
	authorities, ok := mapInfo[string(constants.TokenAuthorities)]
	if !ok {
		return converter.DefaultAuthorities
	}
	return authority.CreateAuthorityList(authorities)
}
