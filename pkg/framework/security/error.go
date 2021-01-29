package security

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"

var (
	// ErrInvalidToken for auth
	ErrInvalidToken = errors.ErrInvalidToken
	// ErrExpiredToken for auth
	ErrExpiredToken = errors.ErrExpiredToken
)
