package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewInt32Handler creates a hander for faking the int32 type
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
