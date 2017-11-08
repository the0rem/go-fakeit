package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
)

// NewStrfmtUUID3Handler creates a hander for faking the strfmt.UUID3 type
func NewStrfmtUUID3Handler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.UUID3",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.UUID3
		NamespaceDNS, _ := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
		item.Scan(uuid.NewV3(NamespaceDNS, fake.DomainName()).String())
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
