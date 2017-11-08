package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

func NewFloat32Handler() *Liar {
	liar := Liar{
		Kind: reflect.Float32,
		Type: "float32",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(r.Float32()))
	}

	return &liar
}
