package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewInt8Handler creates a hander for faking the int8 type
func NewInt8Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Int8,
		Type: "int8",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(int8(r.Int()))
	}

	return &TypeHandler
}
