package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtEmailHandler creates a hander for faking the strfmt.Email type
func NewStrfmtEmailHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.Email",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.Email
		item.Scan(fake.EmailAddress())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
