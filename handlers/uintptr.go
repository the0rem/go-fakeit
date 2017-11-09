package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUIntPtrHandler creates a hander for faking the uintptr type
func NewUIntPtrHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Uintptr,
		Type: "uintptr",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		ptr := uintptr(r.Uint32())
		return reflect.ValueOf(ptr)
	}

	return &TypeHandler
}
