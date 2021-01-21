package provider

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/password"

// BuildPasswordEncoder for inject Encoder
func BuildPasswordEncoder() (password.Encoder, func(), error) {
	encoder := password.NewSha1Encoder()
	return encoder, func() {}, nil
}
