package dto

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/model/enums"

// QueryStatusParams for request params
type QueryStatusParams struct {
	Status enums.CommonStatus `json:"status" form:"status"`
}
