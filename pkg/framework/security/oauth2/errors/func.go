package errors

import (
	"net/http"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
)

// InvalidToken 无效的Token，自定义提示信息
func InvalidToken(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusUnauthorized, TokenInvalid, message)
}

// Forbidden error
func Forbidden(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.Forbidden(message)
}

// Unauthorized error
func Unauthorized(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.Unauthorized(message)
}

// InvalidGrant 无效的授权
func InvalidGrant(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusBadRequest, InvalidGrantCode, message)
}

// InvalidScope 无效的Scope
func InvalidScope(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.BadRequest(message)
}

// InvalidClient 无效的Client
func InvalidClient(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusUnauthorized, InvalidClientCode, message)
}

// InvalidRequest 无效的情况
func InvalidRequest(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.New(http.StatusBadRequest, InvalidRequestCode, message)
}

// InsufficientAuthentication 不充足的认证
func InsufficientAuthentication(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.Unauthorized(message)
}
