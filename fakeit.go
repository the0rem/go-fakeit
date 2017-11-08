package main

import (
	"fmt"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/the0rem/go-fakeit/handlers"
)

const (
	LiarTag = "fakeit"
)

type FakeMaker struct {
	Fakers []*handlers.Liar
}

func (fakeMaker *FakeMaker) CanHandle(field reflect.Value) bool {
	kind := field.Kind()
	fieldType := field.Type().String()
	packagePath := field.Type().PkgPath()

	if !field.CanSet() {
		return false
	}

	for _, handler := range fakeMaker.Fakers {
		if handler.Package == packagePath &&
			handler.Kind == kind &&
			handler.Type == fieldType {
			return true
		}
	}
	return false
}

func (fakeMaker *FakeMaker) GenerateValue(field reflect.Value) {
	kind := field.Kind()
	fieldType := field.Type().String()
	packagePath := field.Type().PkgPath()

	if !field.CanAddr() || !field.CanSet() {
		return
	}

	for _, handler := range fakeMaker.Fakers {
		if handler.Package == packagePath &&
			handler.Kind == kind &&
			handler.Type == fieldType {
			handler.Fill(field, handlers.Tag{})
		}
	}
}

var fakeMaker = FakeMaker{
	Fakers: []*handlers.Liar{
		handlers.NewBoolHandler(),
		handlers.NewFloat32Handler(),
		handlers.NewFloat64Handler(),
		handlers.NewInt16Handler(),
		handlers.NewInt32Handler(),
		handlers.NewInt64Handler(),
		handlers.NewInt8Handler(),
		handlers.NewIntHandler(),
		handlers.NewStringHandler(),
		handlers.NewUInt16Handler(),
		handlers.NewUInt32Handler(),
		handlers.NewUInt64Handler(),
		handlers.NewUInt8Handler(),
		handlers.NewUIntHandler(),
		handlers.NewUIntPtrHandler(),
		handlers.NewStrfmtBase64Handler(),
		handlers.NewStrfmtCreditCardHandler(),
		handlers.NewStrfmtDateHandler(),
		handlers.NewStrfmtDateTimeHandler(),
		handlers.NewStrfmtDurationHandler(),
		handlers.NewStrfmtEmailHandler(),
		handlers.NewStrfmtHexColorHandler(),
		handlers.NewStrfmtHostnameHandler(),
		handlers.NewStrfmtIPv4Handler(),
		handlers.NewStrfmtIPv6Handler(),
		handlers.NewStrfmtISBN10Handler(),
		handlers.NewStrfmtISBN13Handler(),
		handlers.NewStrfmtISBNHandler(),
		handlers.NewStrfmtMACHandler(),
		handlers.NewStrfmtPasswordHandler(),
		handlers.NewStrfmtRGBColorHandler(),
		handlers.NewStrfmtSSNHandler(),
		handlers.NewStrfmtURIHandler(),
		handlers.NewStrfmtUUID3Handler(),
		handlers.NewStrfmtUUID4Handler(),
		handlers.NewStrfmtUUID5Handler(),
		handlers.NewStrfmtUUIDHandler(),
	},
}

type StatusVal int

type BigLebowski struct {
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
	String string

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
	Slice    []BigLebowski

	PtrMap map[string]*BigLebowski
	Map    map[string]BigLebowski

	// Chan chan int

	Interface interface{}

	Ptr       *int
	PtrStruct *BigLebowski
	Struct    BigLebowski

	StatusVal StatusVal
}

type DunWork struct {
	Ptr       *int
	PtrStruct *BigLebowski
	Array     [2]strfmt.Password
	Slice     []BigLebowski
	PtrSlice  []*strfmt.Password
	PtrArray  [2]*strfmt.Password

	Map    map[string]BigLebowski
	MapPtr map[string]*BigLebowski
}

func FillStruct(a interface{}) {
	t := reflect.TypeOf(a)
	valueOf := reflect.ValueOf(a)
	DisectFields(t, valueOf, "")
}

func DisectFields(t reflect.Type, valueOf reflect.Value, logPrefix string) {

	if valueOf.Kind() != reflect.Ptr {
		panic("Aint a pointer: " + valueOf.Kind().String())
	}

	zeVal := reflect.Indirect(valueOf)

	if fakeMaker.CanHandle(zeVal) {
		fakeMaker.GenerateValue(zeVal)
	}

	fmt.Println(logPrefix + "Go to: " + zeVal.Kind().String())

	switch zeVal.Kind() {

	case reflect.Ptr:
		if !zeVal.IsValid() {
			fmt.Println("Aint valid")
			return
		}
		thing := reflect.New(zeVal.Type().Elem())

		if !zeVal.CanSet() {
			return
		}

		zeVal.Set(thing)
		DisectFields(thing.Type(), thing, logPrefix+" -  - ")

	case reflect.Struct:
		for j := 0; j < zeVal.NumField(); j++ {

			field := zeVal.Field(j).Addr()
			fieldName := zeVal.Type().Field(j).Name
			dataType := field.Type().String()
			packagePath := field.Type().PkgPath()
			kind := field.Kind()

			fmt.Printf("%s %s Name: %s  Type: %s  Package: %s\n", logPrefix+" - ", kind, fieldName, dataType, packagePath)

			DisectFields(field.Type(), field, logPrefix+" -  - ")
		}

	case reflect.Map:
		zeVal.Set(reflect.MakeMap(zeVal.Type()))
		keys := zeVal.MapKeys()
		fmt.Println("MAPPPPPP", zeVal.Type().String(), reflect.SliceOf(zeVal.Type().Elem()), zeVal.Len(), keys)

		for _, key := range keys {
			field := zeVal.MapIndex(key).Addr()
			// field := reflect.New(zeVal.Type()).Elem()
			DisectFields(field.Type(), field, logPrefix+" - ")
			zeVal.SetMapIndex(key, field)
			fieldName := key
			dataType := field.Type().String()
			packagePath := field.Type().PkgPath()
			kind := field.Kind()

			fmt.Printf("%s %s Name: %s  Type: %s  Package: %s\n", logPrefix+" - ", kind, fieldName, dataType, packagePath)
			DisectFields(field.Type(), field, logPrefix+" -  - ")
		}

	case reflect.Slice:
		zeVal.Set(reflect.MakeSlice(zeVal.Type(), 1, 1))
		field := zeVal.Index(0).Addr()
		DisectFields(field.Type(), field, logPrefix+" - ")
		// fmt.Println("SLICEEEE", zeVal.Type().String(), reflect.SliceOf(zeVal.Type().Elem()), zeVal.Len())

	case reflect.Array:
		fmt.Printf("%+v\n", zeVal)
		for i := 0; i < zeVal.Len(); i++ {
			field := zeVal.Index(i).Addr()
			DisectFields(field.Type(), field, logPrefix+" - ")
		}

	default:
		// if zeVal.Type().String() == zeVal.Kind().String() {
		if fakeMaker.CanHandle(zeVal) {
			fakeMaker.GenerateValue(zeVal)
		} else {
			fmt.Println("BadSet", zeVal.Type().String(), zeVal.Kind().String())
		}
	}
}

func main() {
	// test := DunWork{}
	test := testStruct{}
	FillStruct(&test)
	fmt.Println("main")
	fmt.Printf("%+v\n", test)
}
