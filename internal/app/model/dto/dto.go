package dto

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/enums"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// QueryCondition 基础查询条件
type QueryCondition struct {
	ID     types.ID           `json:"id" form:"id"`
	IDs    []types.ID         `json:"ids" form:"ids"`
	Status enums.CommonStatus `json:"status" form:"status"`
}
