package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
)

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
