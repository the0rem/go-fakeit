package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUIntPtrHandler creates a hander for faking the uintptr type
func NewUIntPtrHandler() *Liar {
	liar := Liar{
		Kind: reflect.Uintptr,
		Type: "uintptr",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		ptr := uintptr(r.Uint32())
		field.Set(reflect.ValueOf(ptr))
	}

	return &liar
}
