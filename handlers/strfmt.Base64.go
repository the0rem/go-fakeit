package handlers

import (
	"encoding/base64"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtBase64Handler creates a hander for faking the strfmt.Base64 type
func NewStrfmtBase64Handler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.Base64",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.Base64
		item.Scan(base64.StdEncoding.EncodeToString([]byte(fake.Words())))
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
