package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
)

func NewBoolHandler() *Liar {
	liar := Liar{
		Kind: reflect.Bool,
		Type: "bool",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		field.Set(reflect.ValueOf(randomdata.Boolean()))
	}

	return &liar
}
