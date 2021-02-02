package factory

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/factory"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
)

// NewPasswordEncoder for inject Encoder
func NewPasswordEncoder() (password.Encoder, func(), error) {
	encoder := factory.CreateDelegatingPasswordEncoder()
	return encoder, func() {}, nil
}
