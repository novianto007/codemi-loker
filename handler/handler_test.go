package handler

import (
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	key := "key"
	RegisterHandler(key, func(params []string) {})
	if _, ok := routeHandler[key]; !ok {
		t.Fatal("key must exist")
	}
}

func TestHandle(t *testing.T) {
	key := "handle"
	result := ""
	RegisterHandler(key, func(params []string) {
		result = params[0]
	})
	Handle("handle resultval")
	if result != "resultval" {
		t.Fatal("Handle not working")
	}
}
