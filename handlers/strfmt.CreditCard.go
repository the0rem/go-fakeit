package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtCreditCardHandler creates a hander for faking the strfmt.CreditCard type
func NewStrfmtCreditCardHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.CreditCard",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.CreditCard
		item.Scan(fake.CreditCardNum(""))
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
