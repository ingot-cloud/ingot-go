package security

import "github.com/ingot-cloud/ingot-go/internal/app/support/errors"

var (
	// ErrInvalidToken for auth
	ErrInvalidToken = errors.ErrInvalidToken
	// ErrExpiredToken for auth
	ErrExpiredToken = errors.ErrExpiredToken
)
