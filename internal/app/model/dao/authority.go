package dao

import (
	"context"
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"

	"gorm.io/gorm"
)

func getAuthorityDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.Authority))
}

// Authority Dao
type Authority struct {
	DB *gorm.DB
}

// RoleAuthority get
func (a *Authority) RoleAuthority(ctx context.Context, params dto.QueryStatusParams) (*dto.RoleAuthorityResult, error) {

	db := getAuthorityDB(ctx, a.DB)
	if p := params.Status; p != "" {
		db.Where("status = ?", p)
	}

	var list []*dto.RoleAuthority
	query := "gra.role_id, gm_authority.id as authority_id, gm_authority.name, gm_authority.path"
	joins := "inner join gm_role_authority as gra on gra.authority_id = gm_authority.id"
	db.Select(query).Joins(joins).Scan(&list)

	return &dto.RoleAuthorityResult{
		List: list,
	}, nil
}
