package errors

import (
	"fmt"
	"net/http"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/code"
)

// IllegalArgument error
func IllegalArgument(message string) error {
	return New(http.StatusBadRequest, code.IllegalArgument, message)
}

// IllegalOperation error
func IllegalOperation(message string) error {
	return New(http.StatusBadRequest, code.IllegalOperation, message)
}

// Forbidden error
func Forbidden(e error) error {
	return New(http.StatusForbidden, code.Forbidden, e.Error())
}

// Unauthorized error
func Unauthorized(e error) error {
	return New(http.StatusUnauthorized, code.Unauthorized, e.Error())
}

// NoRoute for http resource not found 404
func NoRoute(path string) error {
	return New(http.StatusNotFound, code.NoRoute, fmt.Sprintf("Path [%s] not found", path))
}

// NoMethod for http method not allow 405
func NoMethod(method string) error {
	return New(http.StatusMethodNotAllowed, code.NoMethod, fmt.Sprintf("Method [%s] not allow", method))
}
