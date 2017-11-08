package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

// NewStrfmtUUIDHandler creates a hander for faking the strfmt.UUID type
func NewStrfmtUUIDHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.UUID",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.UUID
		item.Scan(uuid.NewV1().String())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
