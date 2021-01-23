package factory

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"

// CreateDelegatingPasswordEncoder Creates a DelegatingPasswordEncoder with default mappings.
// Additional mappings may be added and the encoding will be updated to conform with best practices.
// However, due to the nature of DelegatingPasswordEncoder the updates should not impact users
func CreateDelegatingPasswordEncoder() password.Encoder {
	idForEncode := "noop"
	encoders := make(map[string]password.Encoder)
	encoders["noop"] = &password.NoopEncoder{}
	encoders["sha1"] = &password.Sha1Encoder{}
	return &password.DelegatingEncoder{
		IDForEncode:         idForEncode,
		IDToPasswordEncoder: nil,
	}
}
