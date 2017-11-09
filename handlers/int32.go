package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewInt32Handler creates a hander for faking the int32 type
func NewInt32Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Int32,
		Type: "int32",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return reflect.ValueOf(r.Int31())
	}

	return &TypeHandler
}
