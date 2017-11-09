package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
)

// NewStrfmtUUID5Handler creates a hander for faking the strfmt.UUID5 type
func NewStrfmtUUID5Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.UUID5",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.UUID5
		NamespaceDNS, _ := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
		item.Scan(uuid.NewV5(NamespaceDNS, fake.DomainName()).String())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
