package config

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"

// HTTPConfigurer http 配置
type HTTPConfigurer interface {
	Configure(*ingot.Router)
	GetAPI() APIConfigurers
}

// APIConfigurer api 接口配置
type APIConfigurer interface {
	Apply(*ingot.Router)
}

// APIConfigurers api 接口配置列表
type APIConfigurers []APIConfigurer
