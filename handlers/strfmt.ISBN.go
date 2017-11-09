package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
)

// NewStrfmtISBNHandler creates a hander for faking the strfmt.ISBN type
func NewStrfmtISBNHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.ISBN",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.ISBN
		item.Scan("")
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
