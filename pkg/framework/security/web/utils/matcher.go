package utils

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/ingot"

// RequestMatcher 请求匹配器
type RequestMatcher func(*ingot.Context) bool
