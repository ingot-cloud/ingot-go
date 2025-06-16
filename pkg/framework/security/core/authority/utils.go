package authority

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// CreateAuthorityList 创建授权列表
func CreateAuthorityList(authorities any) []core.GrantedAuthority {
	if authorities == nil {
		return nil
	}
	switch value := authorities.(type) {
	case string:
		return []core.GrantedAuthority{&SimpleGrantedAuthority{Role: value}}
	case []string:
		result := make([]core.GrantedAuthority, 0, len(value))
		for _, role := range value {
			result = append(result, &SimpleGrantedAuthority{Role: role})
		}
		return result
	default:
		return nil
	}
}

// ToStringArray 授权列表转字符串列表
func ToStringArray(authorities []core.GrantedAuthority) []string {
	result := make([]string, 0, len(authorities))
	for _, item := range authorities {
		result = append(result, item.GetAuthority())
	}
	return result
}
