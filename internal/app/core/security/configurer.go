package security

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

// CustomHTTPConfigurer 安全配置
type CustomHTTPConfigurer struct {
}

// Build 构建安全配置
func (c *CustomHTTPConfigurer) Build() (security.WebSecurityConfigurers, error) {
	var configurers security.WebSecurityConfigurers
	configurers = append(configurers, &config.WebSecurityConfigurerAdapter{
		AdditionalConfigurer: c,
	})
	return configurers, nil
}

// Configure 配置
func (c *CustomHTTPConfigurer) Configure(http security.HTTPSecurityBuilder) error {
	return nil
}
