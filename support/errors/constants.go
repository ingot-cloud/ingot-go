package errors

import (
	"fmt"
	"ingot/support/code"
	"net/http"
)

var (
	// ErrUnknown err
	ErrUnknown = New(http.StatusInternalServerError, code.Unknown, "Unknow")
	// ErrInvalidToken for auth
	ErrInvalidToken = New(http.StatusForbidden, code.TokenInvalid, "Token invalid")
	// ErrExpiredToken for auth
	ErrExpiredToken = New(http.StatusForbidden, code.TokenExpired, "Token expired")
)

// Forbidden error
func Forbidden(e error) error {
	return New(http.StatusForbidden, code.Forbidden, e.Error())
}

// NoRoute for http resource not found 404
func NoRoute(path string) error {
	return New(http.StatusNotFound, code.NoRoute, fmt.Sprintf("Path [%s] not found", path))
}

// NoMethod for http method not allow 405
func NoMethod(method string) error {
	return New(http.StatusMethodNotAllowed, code.NoMethod, fmt.Sprintf("Method [%s] not allow", method))
}
