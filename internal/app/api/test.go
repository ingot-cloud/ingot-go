package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/user"
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

	ingotUser, ok := auth.GetPrincipal().(*user.IngotUser)
	log.Infof("ok=%t, auth=%v", ok, ingotUser)

	var result struct {
		Test string
	}
	result.Test = "aaa"
	return result, nil
}
