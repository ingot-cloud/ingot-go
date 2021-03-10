package api

import (
	"github.com/gin-gonic/gin"
	coreIngot "github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// Test 测试API
type Test struct {
}

// Apply api配置
func (t *Test) Apply(app *coreIngot.Router) {
	router := app.Group("")
	router.POST("/test", t.test)
}

func (t *Test) test(ctx *gin.Context) (interface{}, error) {
	auth := ingot.GetAuthentication(ctx)

	log.Infof("auth=%v", auth)

	var result struct {
		Test string
	}
	result.Test = "aaa"
	return result, nil
}
