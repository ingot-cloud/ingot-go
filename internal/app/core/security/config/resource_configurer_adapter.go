package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/configurer"
)

// ResourceServerAdapter 自定义适配器
type ResourceServerAdapter struct {
	*configurer.ResourceServerConfigurerAdapter
}

// NewResourceServerAdapter 实例化
func NewResourceServerAdapter(parent *configurer.ResourceServerConfigurerAdapter) *ResourceServerAdapter {
	return &ResourceServerAdapter{
		ResourceServerConfigurerAdapter: parent,
	}
}

// WebConfigure Web安全配置
func (adapter *ResourceServerAdapter) WebConfigure(web security.WebSecurityBuilder) error {
	err := adapter.ResourceServerConfigurerAdapter.WebConfigure(web)
	if err != nil {
		return err
	}

	return nil
}

// HTTPConfigure 配置
func (adapter *ResourceServerAdapter) HTTPConfigure(http security.HTTPSecurityBuilder) error {
	err := adapter.ResourceServerConfigurerAdapter.HTTPConfigure(http)
	if err != nil {
		return err
	}

	return nil
}
