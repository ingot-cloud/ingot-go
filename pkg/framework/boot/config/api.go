package config

import (
	"github.com/gin-gonic/gin"
)

// HTTPConfigurer http 配置
type HTTPConfigurer interface {
	Configure(gin.IRouter)
}

// APIConfigurer api 接口配置
type APIConfigurer interface {
	Apply(gin.IRouter)
}

// APIConfigurers api 接口配置列表
type APIConfigurers []APIConfigurer
