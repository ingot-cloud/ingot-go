package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

// OAuth2SecurityConfigurer OAuth2安全配置
type OAuth2SecurityConfigurer struct {
}

// Build 构建安全配置
func (oa *OAuth2SecurityConfigurer) Build() (security.WebSecurityConfigurers, error) {
	var configurers security.WebSecurityConfigurers
	configurers = append(configurers, &config.WebSecurityConfigurerAdapter{
		AdditionalConfigurer: oa,
	})
	return configurers, nil
}

// Configure 配置
func (oa *OAuth2SecurityConfigurer) Configure(http security.HTTPSecurityBuilder) error {
	return nil
}
