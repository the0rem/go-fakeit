package handlers

import (
	"reflect"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtHostnameHandler creates a hander for faking the strfmt.Hostname type
func NewStrfmtHostnameHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.Hostname",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.Hostname
		item.Scan(strings.ToLower(fake.DomainName()))
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
