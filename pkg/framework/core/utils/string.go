package utils

import (
	"strings"
)

// StringCombine 字符串合并
func StringCombine(values ...string) string {
	var builder strings.Builder
	for _, item := range values {
		builder.WriteString(item)
	}
	return builder.String()
}
