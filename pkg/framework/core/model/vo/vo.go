package vo

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/response"

// PageResult 分页结果
type PageResult struct {
	Records    any                  `json:"records"`
	Pagination *response.Pagination `json:"pagination"`
}
