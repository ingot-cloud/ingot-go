package server

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"

	"github.com/gin-gonic/gin"
)

// Router interface
type Router interface {
	Register(app *gin.Engine) error
	Authentication() security.Authentication
}
