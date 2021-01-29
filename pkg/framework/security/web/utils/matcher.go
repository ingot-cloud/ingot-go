package utils

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"

// RequestMatcher 请求匹配器
type RequestMatcher func(*ingot.Context) bool

// AnyRequestMatcher 匹配所有请求
func AnyRequestMatcher(context *ingot.Context) bool {
	return true
}
