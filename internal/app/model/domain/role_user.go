package domain

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
)

// RoleUsers array
type RoleUsers []*RoleUser

// To domain to dto
func (rus RoleUsers) To() dto.RoleUsers {
	target := make(dto.RoleUsers, len(rus))
	for i, item := range rus {
		target[i] = item.To()
	}
	return target
}

// RoleUser 关联
type RoleUser struct {
	RoleID string `gorm:"primary_key:size:20"`
	UserID string `gorm:"primary_key;size:36"`
}

// TableName for user role
func (*RoleUser) TableName() string {
	return "gm_role_user"
}

// To domain to dto
func (r RoleUser) To() *dto.RoleUser {
	target := new(dto.RoleUser)
	utils.Copy(r, target)
	return target
}
