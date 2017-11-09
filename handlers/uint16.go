package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUInt16Handler creates a hander for faking the uint16 type
func NewUInt16Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Uint16,
		Type: "uint16",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(uint16(r.Uint32()))
	}

	return &TypeHandler
}
