package core

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"

	"github.com/gin-gonic/gin"
)

// IRouter interface
type IRouter interface {
	Register(app *gin.Engine) error
	Authentication() security.Authentication
}
