package router

import (
	"github.com/ingot-cloud/ingot-go/internal/app/api"

	"github.com/gin-gonic/gin"
)

// InitAuthRouter login
func InitAuthRouter(routerGroup *gin.RouterGroup, auth *api.Auth) {
	router := routerGroup.Group("auth")
	router.POST("/login", auth.Login)
	router.POST("/logout", auth.Logout)
}
