package handlers

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

// NewStrfmtRGBColorHandler creates a hander for faking the strfmt.RGBColor type
func NewStrfmtRGBColorHandler() *Liar {
	liar := Liar{
		Kind:    reflect.String,
		Type:    "strfmt.RGBColor",
		Package: "github.com/go-openapi/strfmt",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		var item strfmt.RGBColor

		rgbHex := fake.HexColor()
		red, err := hex.DecodeString(rgbHex[0:2])
		if err != nil {
			panic(err)
		}
		green, err := hex.DecodeString(rgbHex[2:4])
		if err != nil {
			panic(err)
		}
		blue, err := hex.DecodeString(rgbHex[4:6])
		if err != nil {
			panic(err)
		}

		item.Scan(fmt.Sprintf("rgb(%s, %s, %s)", red, green, blue))
		field.Set(reflect.ValueOf(item))
	}

	return &liar
}
