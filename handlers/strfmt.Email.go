package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtEmailHandler creates a hander for faking the strfmt.Email type
func NewStrfmtEmailHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.Email",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.Email
		item.Scan(fake.EmailAddress())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
