package token

import (
	"encoding/base64"
	"reflect"
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

type User interface {
	GetName() string
}

type Wang struct {
	Name string
}

func (n *Wang) GetName() string {
	return n.Name
}

func AA() *Wang {
	return nil
}

func TestReflect(t *testing.T) {
	var u User = &Wang{Name: "ceshi"}

	val := reflect.ValueOf(u)
	// val2 := reflect.Indirect(val)
	method := val.MethodByName("GetName")
	t.Log(method)
	result := method.Call(make([]reflect.Value, 0))
	v := result[0].Interface()
	t.Log(v)

	ma := reflect.ValueOf(AA)
	mar := ma.Call(nil)[0].Interface().(*Wang)
	t.Log(mar == nil)
}
