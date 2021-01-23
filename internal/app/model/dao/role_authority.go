package dao

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
	"gorm.io/gorm"
)

func getRoleAuthorityDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.SysRoleAuthority))
}

// RoleAuthority Dao
type RoleAuthority struct {
	DB *gorm.DB
}

// GetRoleAuthorityIDs 获取指定角色的权限ID列表
func (ra *RoleAuthority) GetRoleAuthorityIDs(ctx context.Context, roleID types.ID) (*[]types.ID, error) {
	db := getRoleAuthorityDB(ctx, ra.DB).Where("role_id = ?", roleID)

	var roleAuthorityList []*domain.SysRoleAuthority
	err := db.Scan(&roleAuthorityList).Error

	if err != nil {
		return nil, err
	}

	ids := make([]types.ID, 0, len(roleAuthorityList))
	for _, item := range roleAuthorityList {
		ids = append(ids, item.AuthorityID)
	}

	return &ids, nil
}
