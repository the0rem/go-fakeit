package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUIntHandler creates a hander for faking the uint type
func NewUIntHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Uint,
		Type: "uint",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(uint(r.Uint32()))
	}

	return &TypeHandler
}
