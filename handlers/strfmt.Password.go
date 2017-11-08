package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtPasswordHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.Password",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.Password
		item.Scan(fake.SimplePassword())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
