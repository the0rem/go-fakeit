package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUInt64Handler creates a hander for faking the uint64 type
func NewUInt64Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Uint64,
		Type: "uint64",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(r.Uint64())
	}

	return &TypeHandler
}
