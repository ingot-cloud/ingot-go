package filter

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/ingot"
)

// Filter Web 过滤器
type Filter interface {
	// 过滤器排序
	Order() int
	// 执行过滤器
	DoFilter(context *ingot.Context, chain Chain) error
}

// Chain 过滤链
type Chain interface {
	// 执行过滤链中下一个过滤器
	DoFilter(context *ingot.Context) error
}

// Filters 过滤器列表
type Filters []Filter

func (f Filters) Len() int {
	return len(f)
}

func (f Filters) Less(i, j int) bool {
	return f[i].Order() < f[j].Order()
}

func (f Filters) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// SecurityFilterChain 和请求匹配后执行的过滤器链
type SecurityFilterChain interface {
	// 匹配请求
	Matches(context *ingot.Context) bool
	// 待执行的过滤器
	GetFilters() Filters
}
