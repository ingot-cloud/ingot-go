package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/builders"
)

// BuildWebSecurityFilter 构建 Filter
func BuildWebSecurityFilter(configurers security.WebSecurityConfigurers) (filter.Filter, error) {
	webSecurity := &builders.WebSecurity{}

	for _, configurer := range configurers {
		webSecurity.Apply(configurer)
	}

	return webSecurity.Build()
}
