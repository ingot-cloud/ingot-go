package middleware

import (
	"github.com/ingot-cloud/ingot-go/internal/app/core/security"
	"github.com/ingot-cloud/ingot-go/internal/app/core/wrapper/contextwrapper"
	"github.com/ingot-cloud/ingot-go/internal/app/core/wrapper/ginwrapper"
	"github.com/ingot-cloud/ingot-go/internal/app/support/errors"
	"github.com/ingot-cloud/ingot-go/internal/app/support/response"

	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware for middleware
func UserAuthMiddleware(auth security.Authentication, permits ...PermitFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if Permit(ctx, permits...) {
			ctx.Next()
			return
		}

		user, err := auth.ParseUser(ctx.Request.Context(), ginwrapper.GetToken(ctx))
		if err != nil {
			if err == security.ErrInvalidToken || err == security.ErrExpiredToken {
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
