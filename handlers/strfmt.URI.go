package handlers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtURIHandler creates a hander for faking the strfmt.URI type
func NewStrfmtURIHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.URI",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.URI
		item.Scan(fmt.Sprintf("https://%s/", strings.ToLower(fake.DomainName())))
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
