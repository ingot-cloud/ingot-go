package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

// WebSecurityConfigurers 默认WebSecurityConfigurers
func WebSecurityConfigurers() security.WebSecurityConfigurers {
	var configurers security.WebSecurityConfigurers
	configurers = append(configurers, config.NewWebSecurityConfigurerAdapter(nil, nil))
	return configurers
}
