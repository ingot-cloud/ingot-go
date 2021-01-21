package security

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"

var (
	// ErrInvalidToken for auth
	ErrInvalidToken = errors.ErrInvalidToken
	// ErrExpiredToken for auth
	ErrExpiredToken = errors.ErrExpiredToken
)
