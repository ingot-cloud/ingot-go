package domain

// RoleAuthority 关联
type RoleAuthority struct {
	RoleID      string `gorm:"primary_key;size:20"`
	AuthorityID string `gorm:"primary_key;size:36"`
}

// TableName for RoleAuthority
func (*RoleAuthority) TableName() string {
	return "gm_role_authority"
}
