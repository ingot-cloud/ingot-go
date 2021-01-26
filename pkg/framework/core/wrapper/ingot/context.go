package ingot

import "github.com/gin-gonic/gin"

// Context 自定义 context
type Context struct {
	*gin.Context
}
