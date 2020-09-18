package response

import (
	"ingot/support/code"
	"ingot/support/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result for api request
func Result(ctx *gin.Context, statusCode int, code int, data D, message string) {
	ctx.JSON(statusCode, R{
		code,
		data,
		message,
	})
}

// Ok response struct
func Ok(ctx *gin.Context, data D) {
	Result(ctx, http.StatusOK, code.SUCCESS, data, "Success")
}

// Failure response struct
func Failure(ctx *gin.Context, code int, message string) {
	Result(ctx, http.StatusInternalServerError, code, D{}, message)
}

// FailureWithE response
func FailureWithE(ctx *gin.Context, e *errors.E) {
	Result(ctx, e.StatusCode, e.Code, D{}, e.Message)
}

// FailureWithError response
func FailureWithError(ctx *gin.Context, e error) {
	FailureWithE(ctx, errors.Unpack(e))
}

// Failure500 response
func Failure500(ctx *gin.Context, message string) {
	Failure(ctx, code.InternalServerError, message)
}
