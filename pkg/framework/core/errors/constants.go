package errors

import (
	"net/http"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/code"
)

var (
	// ErrUnknown err
	ErrUnknown = New(http.StatusInternalServerError, code.InternalServerError, "Unknow")
	// ErrUnauthorized for auth
	ErrUnauthorized = New(http.StatusUnauthorized, "", "User unauthorized")
	// ErrUserInvalid for login
	ErrUserInvalid = New(http.StatusBadRequest, "", "Wrong username or password")
	// ErrUserDisabled for login
	ErrUserDisabled = New(http.StatusForbidden, "", "The user has been disabled")
	// ErrIDClockBack clock callback
	ErrIDClockBack = New(http.StatusOK, code.IDClockBack, "Clock callback")
)
