package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

// NewStrfmtUUID4Handler creates a hander for faking the strfmt.UUID4 type
func NewStrfmtUUID4Handler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.UUID4",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.UUID4
		item.Scan(uuid.NewV4().String())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
