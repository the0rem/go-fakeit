package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUInt8Handler creates a hander for faking the uint8 type
func NewUInt8Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Uint8,
		Type: "uint8",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(uint8(r.Uint32()))
	}

	return &TypeHandler
}
