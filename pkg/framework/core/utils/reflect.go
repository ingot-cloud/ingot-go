package utils

import (
	"reflect"
	"strings"
)

// GetType 获取类型
// 返回的 string 为 PkgPath + type
func GetType(target any) string {
	value := reflect.ValueOf(target)
	targetType := reflect.Indirect(value).Type()
	var builder strings.Builder
	builder.WriteString(targetType.PkgPath())
	builder.WriteString("/")
	builder.WriteString(targetType.Name())
	return builder.String()
}
