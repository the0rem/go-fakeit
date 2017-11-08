package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

func NewInt64Handler() *Liar {
	liar := Liar{
		Kind: reflect.Int64,
		Type: "int64",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(r.Int63()))
	}

	return &liar
}
