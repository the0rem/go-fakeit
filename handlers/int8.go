package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewInt8Handler creates a hander for faking the int8 type
func NewInt8Handler() *Liar {
	liar := Liar{
		Kind: reflect.Int8,
		Type: "int8",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(int8(r.Int())))
	}

	return &liar
}
