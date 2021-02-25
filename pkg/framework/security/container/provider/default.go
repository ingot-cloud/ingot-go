package provider

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"

// NilClientdetails 空实现
func NilClientdetails() clientdetails.Service {
	return &clientdetails.NilClientdetails{}
}
