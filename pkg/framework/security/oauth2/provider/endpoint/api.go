package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
)

// OAuth2Api 端点
type OAuth2Api struct {
	TokenEndpoint *TokenEndpoint
}

// Apply api配置
func (a *OAuth2Api) Apply(app *ingot.Router) {
	router := app.Group("oauth")
	router.POST("/token", a.AccessToken)
}

// AccessToken 获取Token
func (a *OAuth2Api) AccessToken(ctx *gin.Context) (interface{}, error) {
	return a.TokenEndpoint.AccessToken(ctx)
}
