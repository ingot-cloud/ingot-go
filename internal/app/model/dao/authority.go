package dao

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"

	"gorm.io/gorm"
)

func getAuthorityDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.SysAuthority))
}

// Authority Dao
type Authority struct {
	DB *gorm.DB
}

// GetAuthoritysWithIDs 根据权限ID获取相应可以使用的权限列表
func (a *Authority) GetAuthoritysWithIDs(ctx context.Context, condition dto.QueryCondition) (*domain.SysAuthoritys, error) {
	db := getAuthorityDB(ctx, a.DB)

	if ids := condition.IDs; len(ids) != 0 {
		db = db.Where("id IN ?", ids)
	}
	if status := condition.Status; status != "" {
		db = db.Where("status = ?", status)
	}

	var list domain.SysAuthoritys
	err := db.Scan(&list).Error

	return &list, err
}

// GetChildWithPID 获取子权限
func (a *Authority) GetChildWithPID(ctx context.Context, condition dto.QueryCondition) (*domain.SysAuthoritys, error) {
	db := getAuthorityDB(ctx, a.DB)

	if status := condition.Status; status != "" {
		db = db.Where("status = ?", status)
	}
	if pid := condition.ID; pid != 0 {
		db = db.Where("pid = ?", pid)
	}

	var list domain.SysAuthoritys
	err := db.Scan(&list).Error

	return &list, err
}
