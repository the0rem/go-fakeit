package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
)

// NewStrfmtISBN10Handler creates a hander for faking the strfmt.ISBN10 type
func NewStrfmtISBN10Handler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.ISBN10",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.ISBN10
		item.Scan("")
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
