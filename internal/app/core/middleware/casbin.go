package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// CasbinMiddleware for middleware
func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, permits ...PermitFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if Permit(ctx, permits...) {
			ctx.Next()
			return
		}

		// path := ctx.Request.URL.Path
		// method := ctx.Request.Method

		// user, ok := ginwrapper.GetUser(ctx)
		// if !ok {
		// 	response.FailureWithError(ctx, errors.ErrUnauthorized)
		// 	ctx.Abort()
		// 	return
		// }

		// if b, err := enforcer.Enforce(user.ID, path, method); err != nil {
		// 	response.FailureWithError(ctx, errors.WithStack(err))
		// 	ctx.Abort()
		// 	return
		// } else if !b {
		// 	response.FailureWithError(ctx, errors.ErrUnauthorized)
		// 	ctx.Abort()
		// 	return
		// }
		ctx.Next()
	}
}
