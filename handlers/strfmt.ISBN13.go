package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
)

// NewStrfmtISBN13Handler creates a hander for faking the strfmt.ISBN13 type
func NewStrfmtISBN13Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.ISBN13",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.ISBN13
		item.Scan("")
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
