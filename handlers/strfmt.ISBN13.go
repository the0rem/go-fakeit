package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
)

func NewStrfmtISBN13Handler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.ISBN13",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.ISBN13
		item.Scan("")
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
