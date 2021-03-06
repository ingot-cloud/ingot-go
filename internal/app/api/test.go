package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
)

// Test 测试API
type Test struct {
}

// Apply api配置
func (t *Test) Apply(app *ingot.Router) {
	router := app.Group("")
	router.POST("/test", t.test)
}

func (t *Test) test(ctx *gin.Context) (interface{}, error) {
	var result struct {
		Test string
	}
	result.Test = "aaa"
	return result, nil
}
