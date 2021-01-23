package dto

// RoleAuthority 角色权限结构
type RoleAuthority struct {
	RoleID      string `json:"roleId"`
	AuthorityID string `json:"authorityId"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Method      string `json:"method"`
}

// RoleAuthorityResult response
type RoleAuthorityResult struct {
	List []*RoleAuthority
}
