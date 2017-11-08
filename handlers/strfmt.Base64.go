package handlers

import (
	"encoding/base64"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtBase64Handler creates a hander for faking the strfmt.Base64 type
func NewStrfmtBase64Handler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.Base64",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.Base64
		item.Scan(base64.StdEncoding.EncodeToString([]byte(fake.Words())))
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
