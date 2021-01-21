package dao

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/core/wrapper/contextwrapper"

	commonDto "github.com/ingot-cloud/ingot-go/pkg/framework/core/model/dto"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/response"

	"gorm.io/gorm"
)

// GetDB db
func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := contextwrapper.GetTrans(ctx)
	if ok && !contextwrapper.GetNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if contextwrapper.GetTransLock(ctx) {
				db = db.Set("gorm:query_option", "FOR UPDATE")
			}
			return db
		}
	}
	return defDB
}

// GetDBWithModel db
func GetDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return GetDB(ctx, defDB).Model(m)
}

// PageScan 分页
func PageScan(ctx context.Context, db *gorm.DB, page commonDto.Pagination, out interface{}) (*response.Pagination, error) {
	return PageScanWithSelect(ctx, db, page, out, "")
}

// PageScanWithSelect 分页查询，设置指定查询结果
func PageScanWithSelect(ctx context.Context, db *gorm.DB, page commonDto.Pagination, out interface{}, query string) (*response.Pagination, error) {
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return nil, err
	} else if count == 0 {
		return &response.Pagination{
			Total:   0,
			Current: 0,
		}, nil
	}

	if query != "" {
		db = db.Select(query)
	}

	current, pageSize := page.Current, page.Size
	if current > 0 && pageSize > 0 {
		db = db.Offset((current - 1) * pageSize).Limit(pageSize)
	} else if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	err = db.Scan(out).Error
	if err != nil {
		return nil, err
	}

	return &response.Pagination{
		Total:   int(count),
		Current: current,
	}, nil
}

// PageFind 分页查询
func PageFind(ctx context.Context, db *gorm.DB, page commonDto.Pagination, out interface{}) (*response.Pagination, error) {
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return nil, err
	} else if count == 0 {
		return &response.Pagination{
			Total:   0,
			Current: 0,
		}, nil
	}

	current, pageSize := page.Current, page.Size
	if current > 0 && pageSize > 0 {
		db = db.Offset((current - 1) * pageSize).Limit(pageSize)
	} else if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	err = db.Find(out).Error
	if err != nil {
		return nil, err
	}

	return &response.Pagination{
		Total:   int(count),
		Current: current,
	}, nil
}
