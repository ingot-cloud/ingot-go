package dto

import "github.com/ingot-cloud/ingot-go/internal/app/model/enums"

// Pagination for request params
type Pagination struct {
	Size    int `json:"size" form:"size"`
	Current int `json:"current" form:"current"`
}

// QueryStatusParams for request params
type QueryStatusParams struct {
	Status enums.CommonStatus `json:"status" form:"status"`
}
