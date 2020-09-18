package router

import (
	"ingot/api"

	"github.com/gin-gonic/gin"
)

// InitTestRouter for test
func InitTestRouter(routerGroup *gin.RouterGroup, test *api.Test) {
	router := routerGroup.Group("test")
	router.GET("hello", test.TestFunc)
}
