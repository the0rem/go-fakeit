package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtIPv6Handler creates a hander for faking the strfmt.IPv6 type
func NewStrfmtIPv6Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.IPv6",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.IPv6
		item.Scan(fake.IPv6())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
