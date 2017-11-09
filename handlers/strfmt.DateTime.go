package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-openapi/strfmt"
)

// NewStrfmtDateTimeHandler creates a hander for faking the strfmt.DateTime type
func NewStrfmtDateTimeHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.DateTime",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.DateTime
		dateTime, _ := strfmt.ParseDateTime(randomdata.FullDate())
		item.Scan(dateTime.String())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
