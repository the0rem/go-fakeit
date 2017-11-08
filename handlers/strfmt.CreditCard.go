package handlers

import (
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtCreditCardHandler creates a hander for faking the strfmt.CreditCard type
func NewStrfmtCreditCardHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.CreditCard",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.CreditCard
		item.Scan(fake.CreditCardNum(""))
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
