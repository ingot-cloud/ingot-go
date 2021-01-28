package security

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"

// DefaultHTTPConfigurer 安全配置
type DefaultHTTPConfigurer struct {
}

// Build 构建安全配置
func (c *DefaultHTTPConfigurer) Build() (config.WebSecurityConfigurers, error) {
	var configurers config.WebSecurityConfigurers
	configurers = append(configurers, &config.WebSecurityConfigurerAdapter{
		AdditionalConfigurer: c,
	})
	return configurers, nil
}

// Configure 配置
func (c *DefaultHTTPConfigurer) Configure(http *config.HTTPSecurity) error {
	return nil
}
