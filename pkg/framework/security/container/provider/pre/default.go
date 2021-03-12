package pre

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// NilClientDetails 空实现
func NilClientDetails() clientdetails.Service {
	return &clientdetails.NilClientdetails{}
}

// NilUserDetailsService 空实现
func NilUserDetailsService() userdetails.Service {
	return &userdetails.NilUserDetailsService{}
}
