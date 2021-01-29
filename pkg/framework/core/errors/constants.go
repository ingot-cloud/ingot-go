package errors

import (
	"net/http"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/code"
)

var (
	// ErrUnknown err
	ErrUnknown = New(http.StatusInternalServerError, code.InternalServerError, "Unknow")
	// ErrUnauthorized for auth
	ErrUnauthorized = New(http.StatusUnauthorized, code.Unauthorized, "User unauthorized")
	// ErrUserInvalid for login
	ErrUserInvalid = New(http.StatusBadRequest, code.UserInvalid, "Wrong username or password")
	// ErrUserDisabled for login
	ErrUserDisabled = New(http.StatusForbidden, code.UserDisabled, "The user has been disabled")
	// ErrUserAppForbidden for login
	ErrUserAppForbidden = New(http.StatusForbidden, code.UserAppForbidden, "No access to the app")
	// ErrIDClockBack clock callback
	ErrIDClockBack = New(http.StatusOK, code.IDClockBack, "Clock callback")
)
