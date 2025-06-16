package ingot

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/response"
)

// HandlerFunc 扩展 gin.HandlerFunc 增加返回值
// 返回第一个参数为响应结构
// 返回第二个参数为异常，那么直接响应该异常
type HandlerFunc func(*gin.Context) (any, error)

func transformHandler(hander HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := hander(ctx)
		if err != nil {
			response.FailureWithError(ctx, err)
			return
		}
		if result != nil {
			response.OK(ctx, result)
			return
		}
		response.OKWithEmpty(ctx)
	}
}

func transformHandlers(handlers ...any) []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, 0, len(handlers))
	for _, handler := range handlers {
		switch value := handler.(type) {
		case HandlerFunc:
			ginHandlers = append(ginHandlers, transformHandler(value))
		case func(*gin.Context) (any, error):
			ginHandlers = append(ginHandlers, transformHandler(value))
		case gin.HandlerFunc:
			ginHandlers = append(ginHandlers, value)
		case func(*gin.Context):
			ginHandlers = append(ginHandlers, value)
		}
	}
	return ginHandlers
}
