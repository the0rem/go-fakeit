package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-openapi/strfmt"
)

func NewStrfmtMACHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.MAC",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.MAC
		item.Scan(randomdata.MacAddress())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
