package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

func NewUInt64Handler() *Liar {
	liar := Liar{
		Kind: reflect.Uint64,
		Type: "uint64",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(r.Uint64()))
	}

	return &liar
}
