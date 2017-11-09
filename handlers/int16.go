package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewInt16Handler creates a hander for faking the int16 type
func NewInt16Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Int16,
		Type: "int16",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(int16(r.Int()))
	}

	return &TypeHandler
}
