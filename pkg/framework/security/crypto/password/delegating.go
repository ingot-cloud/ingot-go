package password

import (
	"strings"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
)

const (
	// PREFIX 前缀
	PREFIX = "{"
	// SUFFIX 后缀
	SUFFIX = "}"
)

// DelegatingEncoder impl
type DelegatingEncoder struct {
	IDForEncode         string
	IDToPasswordEncoder map[string]Encoder
}

// Encode 编码
func (e *DelegatingEncoder) Encode(raw string) (string, error) {
	encoder, ok := e.IDToPasswordEncoder[e.IDForEncode]
	if !ok {
		return "", errors.InternalServer("idForEncode is not found in idToPasswordEncoder")
	}

	result, err := encoder.Encode(raw)

	var builder strings.Builder
	builder.WriteString(PREFIX)
	builder.WriteString(e.IDForEncode)
	builder.WriteString(SUFFIX)
	builder.WriteString(result)

	return builder.String(), err
}

// Matches 验证原始密码和编码后的密码是否相等
func (e *DelegatingEncoder) Matches(raw string, encodedPassword string) (bool, error) {
	if raw == "" && encodedPassword == "" {
		return true, nil
	}
	id := e.extractID(encodedPassword)
	encoder, ok := e.IDToPasswordEncoder[id]
	if !ok {
		return false, errors.InternalServer("id is not found in idToPasswordEncoder")
	}
	password := e.extractEncodedPassword(encodedPassword)
	return encoder.Matches(raw, password)
}

func (e *DelegatingEncoder) extractID(encodedPassword string) string {
	if encodedPassword == "" {
		return ""
	}
	start := strings.Index(encodedPassword, PREFIX)
	if start != 0 {
		return ""
	}
	end := strings.Index(encodedPassword, SUFFIX)
	if end < 0 {
		return ""
	}
	return encodedPassword[start+1 : end]
}

func (e *DelegatingEncoder) extractEncodedPassword(encodedPassword string) string {
	start := strings.Index(encodedPassword, SUFFIX)
	return encodedPassword[start+1:]
}
