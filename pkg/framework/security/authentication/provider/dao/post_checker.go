package dao

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/errors"
)

// DefaultPostAuthenticationChecks 后置检测器
type DefaultPostAuthenticationChecks struct {
}

// NewPostChecker 实例化
func NewPostChecker() *DefaultPostAuthenticationChecks {
	return &DefaultPostAuthenticationChecks{}
}

// Check 检测方法
func (*DefaultPostAuthenticationChecks) Check(user userdetails.UserDetails) error {
	if !user.IsCredentialsNonExpired() {
		return errors.CredentialsExpired("User credentials have expired")
	}
	return nil
}
