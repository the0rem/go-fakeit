package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUInt32Handler creates a hander for faking the uint32 type
func NewUInt32Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Uint32,
		Type: "uint32",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(r.Uint32())
	}

	return &TypeHandler
}
