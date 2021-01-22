package domain

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"

// SysRoleOauthClient 关联
type SysRoleOauthClient struct {
	RoleID   types.ID `gorm:"primary_key;size:20"`
	ClientID types.ID `gorm:"primary_key;size:20"`
}
