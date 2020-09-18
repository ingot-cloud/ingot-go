package api

import (
	"ingot/service"
	"ingot/support/response"

	"github.com/gin-gonic/gin"
)

// Test api
type Test struct {
	Test *service.Test
}

// TestFunc for test
func (t *Test) TestFunc(ctx *gin.Context) {
	t.Test.Test()
	// _ = c.ShouldBindJSON(&dict)
	response.Ok(ctx, response.D{
		"message": "hello world",
	})
}
