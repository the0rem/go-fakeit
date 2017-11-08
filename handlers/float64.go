package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

func NewFloat64Handler() *Liar {
	liar := Liar{
		Kind: reflect.Float64,
		Type: "float64",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(r.Float64()))
	}

	return &liar
}