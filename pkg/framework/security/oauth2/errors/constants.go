package errors

import (
	"net/http"

	coreErrors "github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
)

var (
	// ErrInvalidToken for auth
	ErrInvalidToken = coreErrors.New(http.StatusUnauthorized, TokenInvalid, "Token invalid")
	// ErrExpiredToken for auth
	ErrExpiredToken = coreErrors.New(http.StatusUnauthorized, TokenExpired, "Token expired")
)
