package dao

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/errors"
)

// DefaultPreAuthenticationChecks 默认前置检查器
type DefaultPreAuthenticationChecks struct {
}

// NewPreChecker 实例化
func NewPreChecker() *DefaultPreAuthenticationChecks {
	return &DefaultPreAuthenticationChecks{}
}

func (*DefaultPreAuthenticationChecks) Pre() {}

// Check 检测方法
func (*DefaultPreAuthenticationChecks) Check(user userdetails.UserDetails) error {
	if !user.IsAccountNonLocked() {
		return errors.AccountLock("User account is locked")
	}
	if !user.IsEnabled() {
		return errors.AccountDisabled("User is disabled")
	}
	if !user.IsAccountNonExpired() {
		return errors.AccountExpired("User account has expired")
	}
	return nil
}
