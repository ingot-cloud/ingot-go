package builders

import (
	"sort"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	securityFilter "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// HTTPSecurity http 安全配置
type HTTPSecurity struct {
	requestMatcher utils.RequestMatcher
	filters        filter.Filters
	configurers    []security.HTTPSecurityConfigurer
}

// Build 构建 SecurityFilterChain
func (security *HTTPSecurity) Build() (securityFilter.SecurityFilterChain, error) {
	err := security.configure()
	if err != nil {
		return nil, err
	}

	if security.requestMatcher == nil {
		security.requestMatcher = utils.AnyRequestMatcher
	}

	return security.performBuild(), nil
}

// RequestMatcher 设置请求匹配器
func (security *HTTPSecurity) RequestMatcher(requestMatcher utils.RequestMatcher) {
	security.requestMatcher = requestMatcher
}

// AddFilter 添加 Filter
func (security *HTTPSecurity) AddFilter(filter filter.Filter) {
	security.filters = append(security.filters, filter)
}

// Apply 应用配置
func (security *HTTPSecurity) Apply(configurer security.HTTPSecurityConfigurer) {
	security.configurers = append(security.configurers, configurer)
}

func (security *HTTPSecurity) configure() error {
	// 执行配置
	for _, item := range security.configurers {
		if err := item.Configure(security); err != nil {
			return err
		}
	}
	return nil
}

func (security *HTTPSecurity) performBuild() securityFilter.SecurityFilterChain {
	// 使用升序进行filter排序
	sort.Sort(security.filters)
	return &securityFilter.DefaultSecurityFilterChain{
		RequestMatcher: security.requestMatcher,
		Filters:        security.filters,
	}
}
