package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtHexColorHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.HexColor",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.HexColor
		item.Scan(fake.HexColor())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
