package dto

// RoleUsers array
type RoleUsers []*RoleUser

// RoleUser 关联
type RoleUser struct {
	RoleID string `json:"roleId"`
	UserID string `json:"userId"`
}

// RoleUserQueryParams params
type RoleUserQueryParams struct {
}

// RoleUserQueryResult response
type RoleUserQueryResult struct {
	List RoleUsers
}
