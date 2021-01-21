package response

import (
	"net/http"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/code"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"

	"github.com/gin-gonic/gin"
)

// Result for api request
func Result(ctx *gin.Context, statusCode int, code string, data interface{}, message string) {
	ctx.JSON(statusCode, &R{
		code,
		data,
		message,
	})
}

// PaginationResult for response
func PaginationResult(ctx *gin.Context, data *P) {
	Result(ctx, http.StatusOK, code.SUCCESS, data, "Success")
}

// PaginationWith list and pagination
func PaginationWith(ctx *gin.Context, list interface{}, page *Pagination) {
	PaginationResult(ctx, &P{
		List:       list,
		Pagination: page,
	})
}

// OK response struct
func OK(ctx *gin.Context, data interface{}) {
	Result(ctx, http.StatusOK, code.SUCCESS, data, "Success")
}

// OKWithEmpty response struct
func OKWithEmpty(ctx *gin.Context) {
	Result(ctx, http.StatusOK, code.SUCCESS, &D{}, "Success")
}

// Failure response struct
func Failure(ctx *gin.Context, code string, message string) {
	Result(ctx, http.StatusInternalServerError, code, &D{}, message)
}

// FailureWithE response
func FailureWithE(ctx *gin.Context, e *errors.E) {
	Result(ctx, e.StatusCode, e.Code, &D{}, e.Message)
}

// FailureWithError response
func FailureWithError(ctx *gin.Context, e error) {
	FailureWithE(ctx, errors.Unpack(e))
}

// Failure500 response
func Failure500(ctx *gin.Context, message string) {
	Failure(ctx, code.InternalServerError, message)
}
