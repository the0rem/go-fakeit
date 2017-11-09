package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewFloat32Handler creates a hander for faking the float32 type
func NewFloat32Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Float32,
		Type: "float32",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(r.Float32())
	}

	return &TypeHandler
}
