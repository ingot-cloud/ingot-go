package token

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestBasicToken(t *testing.T) {
	base64Token := "d2ViLWNsb3VkOndlYi1jbG91ZA=="
	raw, _ := base64.StdEncoding.DecodeString(base64Token)
	rawString := string(raw)

	token := strings.Split(rawString, ":")
	t.Log(token[0], token[1])
}
