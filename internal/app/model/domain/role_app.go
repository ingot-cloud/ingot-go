package domain

// RoleApp 关联
type RoleApp struct {
	RoleID string `gorm:"primary_key;size:20"`
	AppID  string `gorm:"primary_key;size:20"`
}

// TableName for RoleApp
func (*RoleApp) TableName() string {
	return "gm_role_app"
}
