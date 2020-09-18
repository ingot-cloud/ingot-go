package middleware

import (
	"ingot/core/security"
	"ingot/core/wrapper/contextwrapper"
	"ingot/core/wrapper/ginwrapper"
	"ingot/support/errors"
	"ingot/support/response"
	"strings"

	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware for middleware
func UserAuthMiddleware(auth security.Authentication, permits ...PermitFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if Permit(ctx, permits...) {
			ctx.Next()
			return
		}

		user, err := auth.ParseUser(ctx.Request.Context(), getToken(ctx))
		if err != nil {
			if err == security.ErrInvalidToken {
				response.FailureWithError(ctx, err)
				ctx.Abort()
				return
			}
			response.FailureWithError(ctx, errors.Forbidden(err))
			ctx.Abort()
			return
		}

		wrapUserAuthContext(ctx, user)
		ctx.Next()
	}
}

func wrapUserAuthContext(c *gin.Context, user *security.User) {
	ginwrapper.SetUser(c, user)
	ctx := contextwrapper.WithUser(c.Request.Context(), user)
	c.Request = c.Request.WithContext(ctx)
}

func getToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}
