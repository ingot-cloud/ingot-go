package errors

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
)

// ProviderNotFound Provider未匹配异常
func ProviderNotFound(args ...string) error {
	message := utils.StringCombine(args...)
	return errors.Unauthorized(message)
}
