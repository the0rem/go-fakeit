package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

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
