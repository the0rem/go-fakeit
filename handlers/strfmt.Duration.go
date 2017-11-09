package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
)

// NewStrfmtDurationHandler creates a hander for faking the strfmt.Duration type
func NewStrfmtDurationHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.Duration",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.Duration
		duration := "1d"
		item.Scan(duration)
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
