package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewInt16Handler creates a hander for faking the int16 type
func NewInt16Handler() *Liar {
	liar := Liar{
		Kind: reflect.Int16,
		Type: "int16",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(int16(r.Int())))
	}

	return &liar
}
