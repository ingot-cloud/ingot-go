package config

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// EnableOAuth2Endpoint 开启端点
func EnableOAuth2Endpoint(securityContainer *container.SecurityContainer, engine *gin.Engine) {
	ingotRouter := ingot.NewRouter(engine.Group(""))
	oauthConfig := securityContainer.AuthorizationServerContainer.TokenEndpointHTTPConfigurer
	oauthConfig.Configure(ingotRouter)
	for _, api := range oauthConfig.GetAPI() {
		api.Apply(ingotRouter)
	}
}
