package provider

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/factory"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
)

// BuildPasswordEncoder for inject Encoder
func BuildPasswordEncoder() (password.Encoder, func(), error) {
	encoder := factory.CreateDelegatingPasswordEncoder()
	return encoder, func() {}, nil
}
