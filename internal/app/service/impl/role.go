package impl

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// Role 服务实现
type Role struct {
}

// GetAllRolesOfUser 获取用户所有角色
func (r *Role) GetAllRolesOfUser(userID, deptID types.ID) []*domain.SysRole {

	return nil
}
