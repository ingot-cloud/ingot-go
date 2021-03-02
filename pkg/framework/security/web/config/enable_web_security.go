package config

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/response"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/builders"
)

var webSecurity *webSecurityConfiguration
var once sync.Once

// EnableWebSecurity 启用 WebSecurity
func EnableWebSecurity(engine *gin.Engine, configurers security.WebSecurityConfigurers) {
	once.Do(func() {
		filter, err := buildWebSecurityFilter(configurers)
		if err != nil {
			panic(err)
		}

		webSecurity = &webSecurityConfiguration{
			Filter: filter,
		}
	})

	engine.Use(webSecurity.middleware())
}

// BuildWebSecurityFilter 构建 Filter
func buildWebSecurityFilter(configurers security.WebSecurityConfigurers) (filter.Filter, error) {
	webSecurity := builders.NewWebSecurity()

	for _, configurer := range configurers {
		webSecurity.Apply(configurer)
	}

	return webSecurity.Build()
}

// 安全配置
type webSecurityConfiguration struct {
	Filter filter.Filter
}

// Middleware 中间件
func (h *webSecurityConfiguration) middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := ingot.NewContext(ctx)
		err := h.Filter.DoFilter(context, internalChain)
		if err != nil {
			response.FailureWithError(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

var internalChain = &emptyChain{}

type emptyChain struct {
}

func (c *emptyChain) DoFilter(context *ingot.Context) error {
	return nil
}
