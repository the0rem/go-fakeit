package handlers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtURIHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.URI",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.URI
		item.Scan(fmt.Sprintf("https://%s/", strings.ToLower(fake.DomainName())))
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
