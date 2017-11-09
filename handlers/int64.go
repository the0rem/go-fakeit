package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewInt64Handler creates a hander for faking the int64 type
func NewInt64Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Int64,
		Type: "int64",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(r.Int63())
	}

	return &TypeHandler
}
