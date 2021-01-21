package dto

// RoleAuthority response
type RoleAuthority struct {
	RoleID      string `json:"roleId"`
	AuthorityID string `json:"authorityId"`
	Name        string `json:"name"`
	Path        string `json:"path"`
}

// RoleAuthorityResult response
type RoleAuthorityResult struct {
	List []*RoleAuthority
}
