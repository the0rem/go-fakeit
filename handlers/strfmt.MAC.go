package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-openapi/strfmt"
)

// NewStrfmtMACHandler creates a hander for faking the strfmt.MAC type
func NewStrfmtMACHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.MAC",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.MAC
		item.Scan(randomdata.MacAddress())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
