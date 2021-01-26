package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/filter"
)

// WebSecurityConfigurer Web security 配置
type WebSecurityConfigurer interface {
}

// HTTPSecurityConfigurer HTTP security 配置
type HTTPSecurityConfigurer interface {
}

// WebSecurityBuilder 构造器
type WebSecurityBuilder interface {
	Build() filter.Filter
}

// HTTPSecurityBuilder 构造器
type HTTPSecurityBuilder interface {
	Build() filter.SecurityFilterChain
}