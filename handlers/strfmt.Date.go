package handlers

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-openapi/strfmt"
)

// NewStrfmtDateHandler creates a hander for faking the strfmt.Date type
func NewStrfmtDateHandler() *TypeHandler {
	TypeHandler := TypeHandler{
		Kind:    reflect.String,
		Type:    "strfmt.Date",
		Package: "github.com/go-openapi/strfmt",
	}

	TypeHandler.GetDefaultFaker = func() reflect.Value {
		var item strfmt.Date
		dateTime, _ := strfmt.ParseDateTime(randomdata.FullDate())
		item.Scan(dateTime.String())
		return reflect.ValueOf(item)
	}

	return &TypeHandler
}
