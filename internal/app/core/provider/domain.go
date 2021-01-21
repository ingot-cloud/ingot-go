package provider

import "github.com/ingot-cloud/ingot-go/internal/app/model/domain"

// Get all domain
func getDomain() []interface{} {
	return []interface{}{
		new(domain.User),
		new(domain.Authority),
		new(domain.Role),
		new(domain.RoleUser),
		new(domain.RoleAuthority),
		new(domain.RoleApp),
	}
}
