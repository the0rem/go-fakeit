package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewIntHandler creates a hander for faking the int type
func NewIntHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Int,
		Type: "int",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(r.Int())
	}

	return &TypeHandler
}
