package errors

import (
	"fmt"
	"net/http"

	"github.com/ingot-cloud/ingot-go/internal/app/support/code"
)

var (
	// ErrUnknown err
	ErrUnknown = New(http.StatusInternalServerError, code.Unknown, "Unknow")
	// ErrUnauthorized for auth
	ErrUnauthorized = New(http.StatusUnauthorized, code.Unauthorized, "User unauthorized")
	// ErrInvalidToken for auth
	ErrInvalidToken = New(http.StatusForbidden, code.TokenInvalid, "Token invalid")
	// ErrExpiredToken for auth
	ErrExpiredToken = New(http.StatusForbidden, code.TokenExpired, "Token expired")
	// ErrUserInvalid for login
	ErrUserInvalid = New(http.StatusForbidden, code.UserInvalid, "Wrong username or password")
	// ErrUserDisabled for login
	ErrUserDisabled = New(http.StatusForbidden, code.UserDisabled, "The user has been disabled")
	// ErrUserAppForbidden for login
	ErrUserAppForbidden = New(http.StatusForbidden, code.UserAppForbidden, "No access to the app")
	// ErrUserEnterpriseDisabled for login
	ErrUserEnterpriseDisabled = New(http.StatusForbidden, code.UserEnterpriseDisabled, "The enterprise has been disabled")
	// ErrIDClockBack clock callback
	ErrIDClockBack = New(http.StatusOK, code.IDClockBack, "Clock callback")
)

// IllegalArgument error
func IllegalArgument(message string) error {
	return New(http.StatusOK, code.IllegalArgument, message)
}

// IllegalOperation error
func IllegalOperation(message string) error {
	return New(http.StatusOK, code.IllegalOperation, message)
}

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
