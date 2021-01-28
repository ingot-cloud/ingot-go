package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/builders"
)

// WebSecurityConfigurerAdapter 安全配置适配器
type WebSecurityConfigurerAdapter struct {
	AdditionalConfigurer security.HTTPSecurityConfigurer
}

// Configure Web安全配置
func (adapter *WebSecurityConfigurerAdapter) Configure(web security.WebSecurityBuilder) error {
	http, err := adapter.getHTTP()
	if err != nil {
		return err
	}

	web.AddSecurityFilterChainBuilder(http)

	return nil
}

func (adapter *WebSecurityConfigurerAdapter) getHTTP() (security.HTTPSecurityBuilder, error) {

	http := &builders.HTTPSecurity{}

	err := adapter.applyDefaultConfiguration(http)
	if err != nil {
		return nil, err
	}

	err = adapter.AdditionalConfigurer.Configure(http)

	return http, err
}

func (adapter *WebSecurityConfigurerAdapter) applyDefaultConfiguration(http security.HTTPSecurityBuilder) error {
	// 应用默认配置
	// http.A()  http.B()
	return nil
}
