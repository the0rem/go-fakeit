package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewFloat64Handler creates a hander for faking the float64 type
func NewFloat64Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Float64,
		Type: "float64",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(r.Float64())
	}

	return &TypeHandler
}
