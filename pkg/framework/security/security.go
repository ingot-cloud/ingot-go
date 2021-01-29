package security

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/response"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// Handler 安全处理器
type Handler struct {
	Filter filter.Filter
}

// Middleware 中间件
func (h *Handler) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := &ingot.Context{
			Context: ctx,
		}
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
