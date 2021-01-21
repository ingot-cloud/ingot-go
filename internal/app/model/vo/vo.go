package vo

import "github.com/ingot-cloud/ingot-go/internal/app/support/response"

// PageResult 分页结果
type PageResult struct {
	Records    interface{}          `json:"records"`
	Pagination *response.Pagination `json:"pagination"`
}
