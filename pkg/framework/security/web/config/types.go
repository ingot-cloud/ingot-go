package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/filter"
)

// WebSecurityConfigurer Web security 配置
type WebSecurityConfigurer interface {
	Configure(*WebSecurity) error
}

// HTTPSecurityConfigurer HTTP security 配置
type HTTPSecurityConfigurer interface {
	Configure(*HTTPSecurity) error
}

// WebSecurityBuilder 构造器
type WebSecurityBuilder interface {
	Build() (filter.Filter, error)
}

// HTTPSecurityBuilder 构造器
type HTTPSecurityBuilder interface {
	Build() (filter.SecurityFilterChain, error)
}

// WebSecurityConfigurers 定义 Web Security 配置列表
type WebSecurityConfigurers []WebSecurityConfigurer
