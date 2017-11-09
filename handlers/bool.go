package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
)

// NewBoolHandler creates a hander for faking the bool type
func NewBoolHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind: reflect.Bool,
		Type: "bool",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		return reflect.ValueOf(randomdata.Boolean())
	}

	return &TypeHandler
}
