package builders

import (
	"sort"

	coreUtils "github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	securityFilter "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// HTTPSecurity http 安全配置
type HTTPSecurity struct {
	requestMatcher utils.RequestMatcher
	filters        map[string]filter.Filter
	configurers    map[string]security.HTTPSecurityConfigurer
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
	typeStr := coreUtils.GetType(filter)
	if _, ok := security.filters[typeStr]; !ok {
		security.filters[typeStr] = filter
	}
}

// Apply 应用配置
func (security *HTTPSecurity) Apply(configurer security.HTTPSecurityConfigurer) {
	typeStr := coreUtils.GetType(configurer)
	if _, ok := security.configurers[typeStr]; !ok {
		security.configurers[typeStr] = configurer
	}
}

// HTTPBasic 开启 basic 验证
func (security *HTTPSecurity) HTTPBasic() {

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
	size := len(security.filters)
	filters := make(filter.Filters, 0, size)
	for _, filter := range security.filters {
		filters = append(filters, filter)
	}

	// 使用升序进行filter排序
	sort.Sort(filters)
	return &securityFilter.DefaultSecurityFilterChain{
		RequestMatcher: security.requestMatcher,
		Filters:        filters,
	}
}
