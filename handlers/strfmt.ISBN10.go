package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
)

// NewStrfmtISBN10Handler creates a hander for faking the strfmt.ISBN10 type
func NewStrfmtISBN10Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.ISBN10",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.ISBN10
		item.Scan("")
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
