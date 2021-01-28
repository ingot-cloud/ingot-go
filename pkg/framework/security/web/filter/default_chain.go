package filter

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// DefaultSecurityFilterChain 默认安全过滤链
type DefaultSecurityFilterChain struct {
	RequestMatcher utils.RequestMatcher
	Filters        filter.Filters
}

// Matches 匹配请求
func (c *DefaultSecurityFilterChain) Matches(context *ingot.Context) bool {
	return c.RequestMatcher(context)
}

// GetFilters 待执行的过滤器
func (c *DefaultSecurityFilterChain) GetFilters() filter.Filters {
	return c.Filters
}
