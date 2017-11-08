package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUInt8Handler creates a hander for faking the uint8 type
func NewUInt8Handler() *Liar {
	liar := Liar{
		Kind: reflect.Uint8,
		Type: "uint8",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(uint8(r.Uint32())))
	}

	return &liar
}
