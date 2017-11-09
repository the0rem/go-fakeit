package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

// NewStrfmtUUID4Handler creates a hander for faking the strfmt.UUID4 type
func NewStrfmtUUID4Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.UUID4",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.UUID4
		item.Scan(uuid.NewV4().String())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
