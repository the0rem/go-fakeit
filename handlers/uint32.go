package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUInt32Handler creates a hander for faking the uint32 type
func NewUInt32Handler() *Liar {
	liar := Liar{
		Kind: reflect.Uint32,
		Type: "uint32",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(r.Uint32()))
	}

	return &liar
}
