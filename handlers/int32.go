package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

func NewInt32Handler() *Liar {
	liar := Liar{
		Kind: reflect.Int32,
		Type: "int32",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(r.Int31()))
	}

	return &liar
}
