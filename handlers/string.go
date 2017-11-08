package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
)

// NewStringHandler creates a hander for faking the string type
func NewStringHandler() *Liar {
	liar := Liar{
		Kind: reflect.String,
		Type: "string",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		field.Set(reflect.ValueOf(randomdata.SillyName()))
	}

	return &liar
}
