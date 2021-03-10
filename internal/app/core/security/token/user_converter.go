package token

import (
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/user"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/authority"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
)

// IngotUserAuthenticationConverter 自定义
type IngotUserAuthenticationConverter struct {
}

// ConvertUserAuthentication 在身份验证信息中提取访问令牌使用的信息
func (converter *IngotUserAuthenticationConverter) ConvertUserAuthentication(auth core.Authentication) (map[string]interface{}, error) {
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
func (converter *IngotUserAuthenticationConverter) ExtractAuthentication(mapInfo map[string]interface{}) (core.Authentication, error) {
	if principal, ok := mapInfo[string(constants.TokenUsername)]; ok {
		authorities := converter.getAuthorities(mapInfo)
		username, _ := principal.(string)
		userID := types.NewIDFrom(mapInfo[EnhancerUserID])
		deptID := types.NewIDFrom(mapInfo[EnhancerDeptID])
		tenantID := types.NewIDFrom(mapInfo[EnhancerTenantID])
		authType, _ := mapInfo[EnhancerAuthType].(string)

		user := user.NewPermitIngotUser(userID, deptID, tenantID, authType, username, "N/A", authorities)
		return authentication.NewAuthenticatedUsernamePasswordAuthToken(user, "N/A", authorities), nil
	}
	return nil, nil
}

func (converter *IngotUserAuthenticationConverter) getAuthorities(mapInfo map[string]interface{}) []core.GrantedAuthority {
	authorities, ok := mapInfo[string(constants.TokenAuthorities)]
	if !ok {
		return nil
	}
	return authority.CreateAuthorityList(authorities)
}
