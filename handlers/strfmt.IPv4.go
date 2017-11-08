package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtIPv4Handler creates a hander for faking the strfmt.IPv4 type
func NewStrfmtIPv4Handler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.IPv4",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.IPv4
		item.Scan(fake.IPv4())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
