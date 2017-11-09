package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

// NewStrfmtUUIDHandler creates a hander for faking the strfmt.UUID type
func NewStrfmtUUIDHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.UUID",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.UUID
		item.Scan(uuid.NewV1().String())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
