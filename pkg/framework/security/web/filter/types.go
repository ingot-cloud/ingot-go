package filter

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/ingot"
)

// SecurityFilterChain 和请求匹配后执行的过滤器链
type SecurityFilterChain interface {
	// 匹配请求
	Matches(context *ingot.Context) bool
	// 待执行的过滤器
	GetFilters() filter.Filters
}
