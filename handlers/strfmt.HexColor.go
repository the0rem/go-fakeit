package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtHexColorHandler creates a hander for faking the strfmt.HexColor type
func NewStrfmtHexColorHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.HexColor",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.HexColor
		item.Scan(fake.HexColor())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
