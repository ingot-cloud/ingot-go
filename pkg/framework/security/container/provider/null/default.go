package null

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// ClientDetails 空实现
func ClientDetails() clientdetails.Service {
	return &clientdetails.NilClientdetails{}
}

// UserDetailsService 空实现
func UserDetailsService() userdetails.Service {
	return &userdetails.NilUserDetailsService{}
}
