package handlers

import (
	"reflect"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtHostnameHandler creates a hander for faking the strfmt.Hostname type
func NewStrfmtHostnameHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.Hostname",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.Hostname
		item.Scan(strings.ToLower(fake.DomainName()))
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
