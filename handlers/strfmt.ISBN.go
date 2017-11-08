package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
)

func NewStrfmtISBNHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.ISBN",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.ISBN
		item.Scan("")
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
