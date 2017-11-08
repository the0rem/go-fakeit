package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtIPv6Handler creates a hander for faking the strfmt.IPv6 type
func NewStrfmtIPv6Handler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.IPv6",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.IPv6
		item.Scan(fake.IPv6())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
