package dao

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"

	"gorm.io/gorm"
)

func getRoleUserDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.SysRoleUser))
}

// RoleUser DAO
type RoleUser struct {
	DB *gorm.DB
}

// GetUserRoleIDs 获取用户的角色
func (ru *RoleUser) GetUserRoleIDs(ctx context.Context, userID types.ID) (*[]types.ID, error) {
	db := getRoleUserDB(ctx, ru.DB).Where("user_id = ?", userID)

	var list []*domain.SysRoleUser
	err := db.Scan(&list).Error

	if err != nil {
		return nil, err
	}

	ids := make([]types.ID, 0, len(list))
	for _, item := range list {
		ids = append(ids, item.RoleID)
	}

	// todo 需要获取用户所属部门的角色，和当前角色进行合并

	return &ids, nil
}
