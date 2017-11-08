package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-openapi/strfmt"
)

func NewStrfmtDateHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.Date",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.Date
		dateTime, _ := strfmt.ParseDateTime(randomdata.FullDate())
		item.Scan(dateTime.String())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
