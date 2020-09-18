package core

import (
	"ingot/core/security"

	"github.com/gin-gonic/gin"
)

// IRouter interface
type IRouter interface {
	Register(app *gin.Engine) error
	Authentication() security.Authentication
}
