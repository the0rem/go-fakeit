package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewIntHandler creates a hander for faking the int type
func NewIntHandler() *Liar {
	liar := Liar{
		Kind: reflect.Int,
		Type: "int",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(r.Int()))
	}

	return &liar
}
