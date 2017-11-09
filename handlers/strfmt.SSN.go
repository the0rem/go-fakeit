package handlers

import (
	"fmt"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtSSNHandler creates a hander for faking the strfmt.SSN type
func NewStrfmtSSNHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.SSN",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.SSN
		item.Scan(fmt.Sprintf("%s-%s-%s", fake.DigitsN(3), fake.DigitsN(2), fake.DigitsN(4)))
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
