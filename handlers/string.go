package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
)

// NewStringHandler creates a hander for faking the string type
func NewStringHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.String,
		Type: "string",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		return reflect.ValueOf(randomdata.SillyName())
	}

	return &TypeHandler
}
