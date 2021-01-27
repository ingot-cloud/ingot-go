package config

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"

// BuildWebSecurityFilter 构建 Filter
func BuildWebSecurityFilter(configurers WebSecurityConfigurers) (filter.Filter, error) {
	webSecurity := &WebSecurity{}

	for _, configurer := range configurers {
		webSecurity.Apply(configurer)
	}

	return webSecurity.Build()
}
