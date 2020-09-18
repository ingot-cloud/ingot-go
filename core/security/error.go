package security

import "ingot/support/errors"

var (
	// ErrInvalidToken for auth
	ErrInvalidToken = errors.ErrInvalidToken
	// ErrExpiredToken for auth
	ErrExpiredToken = errors.ErrExpiredToken
)
