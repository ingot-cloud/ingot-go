package service

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// Role 角色服务
type Role interface {
	// 获取用户所有角色
	GetAllRolesOfUser(userID, deptID types.ID) []*domain.SysRole
}
