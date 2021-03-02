package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/endpoint"
)

// OAuth2 端点
type OAuth2 struct {
	TokenEndpoint *endpoint.TokenEndpoint
}

// Apply api配置
func (a *OAuth2) Apply(app *ingot.Router) {
	router := app.Group("oauth")
	router.POST("/token", a.AccessToken)
}

// AccessToken 获取Token
func (a *OAuth2) AccessToken(ctx *gin.Context) (interface{}, error) {
	return a.TokenEndpoint.AccessToken(ctx)
}
