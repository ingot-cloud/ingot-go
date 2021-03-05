package dto

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/enums"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// UserDetailsDto 用户详情
type UserDetailsDto struct {
	ClientID   string
	TenantID   string
	Mode       enums.UserDetailsModeEnum
	UniqueCode string
}

// UserAuthDetails 详情
type UserAuthDetails struct {
	ID       types.ID
	DeptID   types.ID
	TenantID types.ID
	Username string
	Password string
	Status   enums.UserStatusEnum
	AuthType string
	Roles    []string
}
