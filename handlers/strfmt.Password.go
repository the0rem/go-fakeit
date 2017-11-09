package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtPasswordHandler creates a hander for faking the strfmt.Password type
func NewStrfmtPasswordHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.Password",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.Password
		item.Scan(fake.SimplePassword())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
