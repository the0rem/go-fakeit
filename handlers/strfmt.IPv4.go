package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtIPv4Handler creates a hander for faking the strfmt.IPv4 type
func NewStrfmtIPv4Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.IPv4",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.IPv4
		item.Scan(fake.IPv4())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
