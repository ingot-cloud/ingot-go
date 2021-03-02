package api

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"

// HTTPConfigurer http 配置
type HTTPConfigurer interface {
	Configure(*ingot.Router)
	GetAPI() Configurers
}

// Configurer api 接口配置
type Configurer interface {
	Apply(*ingot.Router)
}

// Configurers api 接口配置列表
type Configurers []Configurer
