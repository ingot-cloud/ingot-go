package errors

import (
	"net/http"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
)

// ProviderNotFound Provider未匹配异常
func ProviderNotFound(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.Unauthorized(message)
}

// UsernameNotFound 用户未找到
func UsernameNotFound(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusUnauthorized, BadCredentialsCode, message)
}

// BadCredentials 错误凭证
func BadCredentials(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusUnauthorized, BadCredentialsCode, message)
}

// AccountLock 账户已锁定
func AccountLock(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusUnauthorized, AccountLockCode, message)
}

// AccountDisabled 账户禁用
func AccountDisabled(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusUnauthorized, AccountDisabledCode, message)
}

// AccountExpired 账户过期
func AccountExpired(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusUnauthorized, AccountExpiredCode, message)
}

// CredentialsExpired 凭证过期
func CredentialsExpired(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusUnauthorized, CredentialsExpiredCode, message)
}
