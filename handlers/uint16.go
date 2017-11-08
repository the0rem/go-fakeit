package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

func NewUInt16Handler() *Liar {
	liar := Liar{
		Kind: reflect.Uint16,
		Type: "uint16",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(uint16(r.Uint32())))
	}

	return &liar
}
