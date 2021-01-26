package config

import (
	"sort"

	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// HTTPSecurity http 安全配置
type HTTPSecurity struct {
	requestMatcher utils.RequestMatcher
	filters        filter.Filters
}

// Build 构建 SecurityFilterChain
func (security *HTTPSecurity) Build() filter.SecurityFilterChain {
	// 升序排序
	sort.Sort(security.filters)
	return &filter.DefaultSecurityFilterChain{
		RequestMatcher: security.requestMatcher,
		Filters:        security.filters,
	}
}

// RequestMatcher 设置请求匹配器
func (security *HTTPSecurity) RequestMatcher(requestMatcher utils.RequestMatcher) {
	security.requestMatcher = requestMatcher
}

// AddFilter 添加 Filter
func (security *HTTPSecurity) AddFilter(filter filter.Filter) {
	security.filters = append(security.filters, filter)
}
