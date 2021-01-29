package errors

import (
	"net/http"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
)

// InvalidToken 无效的Token，自定义提示信息
func InvalidToken(message string) error {
	return errors.New(http.StatusUnauthorized, TokenInvalid, message)
}
