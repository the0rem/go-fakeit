package main

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/the0rem/go-fakeit"
)

type statusVal int

type bigLebowski struct {
	Name  string
	Age   int
	Derps strfmt.CreditCard
	Slice []strfmt.Password
}

type testStruct struct {
	Date       strfmt.Date
	DateTime   strfmt.DateTime
	Duration   strfmt.Duration
	Email      strfmt.Email
	Base64     strfmt.Base64
	URI        strfmt.URI
	Hostname   strfmt.Hostname
	IPv4       strfmt.IPv4
	IPv6       strfmt.IPv6
	MAC        strfmt.MAC
	UUID       strfmt.UUID
	UUID3      strfmt.UUID3
	UUID4      strfmt.UUID4
	UUID5      strfmt.UUID5
	ISBN       strfmt.ISBN
	ISBN10     strfmt.ISBN10
	ISBN13     strfmt.ISBN13
	CreditCard strfmt.CreditCard
	SSN        strfmt.SSN
	HexColor   strfmt.HexColor
	RGBColor   strfmt.RGBColor
	Password   strfmt.Password

	Bool   bool
	String string `fakeit:"password,min=25,max=30"`

	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Uintptr    uintptr
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128

	Array    [2]strfmt.Password
	PtrArray [2]*strfmt.Password
	PtrSlice []*strfmt.Password
	Slice    []bigLebowski

	PtrMap map[string]*bigLebowski
	Map    map[string]bigLebowski

	// Chan chan int

	Interface interface{}

	Ptr       *int
	PtrStruct *bigLebowski
	Struct    bigLebowski

	statusVal statusVal
}

type dunWork struct {
	Ptr       *int
	PtrStruct *bigLebowski
	Array     [2]strfmt.Password
	Slice     []bigLebowski
	PtrSlice  []*strfmt.Password
	PtrArray  [2]*strfmt.Password

	Map    map[string]bigLebowski
	MapPtr map[string]*bigLebowski
}

func main() {
	// test := dunWork{}
	test := testStruct{}
	fakeit.FakeIt(&test)
	fmt.Println("main")
	fmt.Printf("%+v\n", test)
}
