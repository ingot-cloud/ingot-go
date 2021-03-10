package config

import (
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/configurer"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// ResourceServerAdapter 自定义适配器
type ResourceServerAdapter struct {
	*configurer.ResourceServerConfigurerAdapter
	ignoredRequestMatcher utils.RequestMatcher
}

// NewResourceServerAdapter 实例化
func NewResourceServerAdapter(parent *configurer.ResourceServerConfigurerAdapter, ignoreMatcher utils.RequestMatcher) *ResourceServerAdapter {
	result := &ResourceServerAdapter{}
	result.ResourceServerConfigurerAdapter = parent
	result.AdditionalHTTPSecurityConfigurer = result
	result.ignoredRequestMatcher = ignoreMatcher
	return result
}

// WebConfigure Web安全配置
func (adapter *ResourceServerAdapter) WebConfigure(web security.WebSecurityBuilder) error {
	err := adapter.ResourceServerConfigurerAdapter.WebConfigure(web)
	if err != nil {
		return err
	}

	if adapter.ignoredRequestMatcher != nil {
		web.AddIgnoreRequestMatcher(adapter.ignoredRequestMatcher)
	}
	return nil
}

// HTTPConfigure 配置
func (adapter *ResourceServerAdapter) HTTPConfigure(http security.HTTPSecurityBuilder) error {
	err := adapter.ResourceServerConfigurerAdapter.HTTPConfigure(http)
	if err != nil {
		return err
	}

	http.AddFilter(filter.NewTenantFilter())

	return nil
}
