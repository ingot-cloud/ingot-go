package security

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

// Build 构建安全配置
func Build() (security.WebSecurityConfigurers, error) {
	var configurers security.WebSecurityConfigurers
	configurers = append(configurers, &config.WebSecurityConfigurerAdapter{
		AdditionalConfigurer: &CustomHTTPConfigurer{},
	})
	return configurers, nil
}

// CustomHTTPConfigurer 安全配置
type CustomHTTPConfigurer struct {
}

// Configure 配置
func (c *CustomHTTPConfigurer) Configure(http security.HTTPSecurityBuilder) error {
	return nil
}
