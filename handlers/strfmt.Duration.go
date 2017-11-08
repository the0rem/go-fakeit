package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
)

func NewStrfmtDurationHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.Duration",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.Duration
		duration := "1d"
		if args.HasIndex(0) {
			duration = args.GetIndex(0)
		}
		item.Scan(duration)
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
