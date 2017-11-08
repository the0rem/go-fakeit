package handlers

import (
	"fmt"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtSSNHandler creates a hander for faking the strfmt.SSN type
func NewStrfmtSSNHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.SSN",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.SSN
		item.Scan(fmt.Sprintf("%s-%s-%s", fake.DigitsN(3), fake.DigitsN(2), fake.DigitsN(4)))
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
