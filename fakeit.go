package fakeit

import (
	"fmt"
	"reflect"

	"github.com/the0rem/go-fakeit/fakers"
	"github.com/the0rem/go-fakeit/handlers"
)

// FakeMaker handles filling variables with fake data
type FakeMaker struct {
	Fakers       map[string]func(args ...interface{}) interface{}
	TypeHandlers []*handlers.TypeHandler
}

// CanHandle determines whether a reflect.Value field can be faked by the provided handlers
func (fakeMaker *FakeMaker) CanHandle(field reflect.Value) bool {
	kind := field.Kind()
	fieldType := field.Type().String()
	packagePath := field.Type().PkgPath()

	if !field.CanSet() {
		return false
	}

	for _, handler := range fakeMaker.TypeHandlers {
		if handler.Package == packagePath &&
			handler.Kind == kind &&
			handler.Type == fieldType {
			return true
		}
	}
	return false
}

// GenerateValue sets the value of the fiven reflect.Value with the appopriate type
func (fakeMaker *FakeMaker) GenerateValue(field reflect.Value, tag *handlers.Tag) {
	kind := field.Kind()
	fieldType := field.Type().String()
	packagePath := field.Type().PkgPath()

	if !field.CanAddr() || !field.CanSet() {
		return
	}

	for _, handler := range fakeMaker.TypeHandlers {
		if handler.Package == packagePath &&
			handler.Kind == kind &&
			handler.Type == fieldType {
			handler.FakeIt(field, tag)
		}
	}
}

var fakeMaker = FakeMaker{
	Fakers: fakers.NewFakers(),
	TypeHandlers: []*handlers.TypeHandler{
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

// FakeIt takes a struct and populates it with fake data.
func FakeIt(a interface{}) {
	t := reflect.TypeOf(a)
	valueOf := reflect.ValueOf(a)
	DisectFields(t, valueOf, "", handlers.NewTagHandler(""))
}

// DisectFields takes the reflect properties of an interface and recursively populates the
// associated proerties using the appropriate fakers.
func DisectFields(t reflect.Type, valueOf reflect.Value, logPrefix string, tagHandler *handlers.Tag) {

	if valueOf.Kind() != reflect.Ptr {
		panic("Aint a pointer: " + valueOf.Kind().String())
	}

	zeVal := reflect.Indirect(valueOf)

	if fakeMaker.CanHandle(zeVal) {
		fakeMaker.GenerateValue(zeVal, tagHandler)
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
		DisectFields(thing.Type(), thing, logPrefix+" -  - ", handlers.NewTagHandler(""))

	case reflect.Struct:
		for j := 0; j < zeVal.NumField(); j++ {

			field := zeVal.Field(j).Addr()
			fieldName := zeVal.Type().Field(j).Name

			dataType := field.Type().String()
			packagePath := field.Type().PkgPath()
			kind := field.Kind()

			fmt.Printf("%s %s Name: %s Type: %s  Package: %s\n", logPrefix+" - ", kind, fieldName, dataType, packagePath)

			tag := handlers.NewTagHandler(string(zeVal.Type().Field(j).Tag))
			DisectFields(field.Type(), field, logPrefix+" -  - ", tag)
		}

	case reflect.Map:
		zeVal.Set(reflect.MakeMap(zeVal.Type()))
		keys := zeVal.MapKeys()
		fmt.Println("MAPPPPPP", zeVal.Type().String(), reflect.SliceOf(zeVal.Type().Elem()), zeVal.Len(), keys)

		for _, key := range keys {
			field := zeVal.MapIndex(key).Addr()
			// field := reflect.New(zeVal.Type()).Elem()
			DisectFields(field.Type(), field, logPrefix+" - ", handlers.NewTagHandler(""))
			zeVal.SetMapIndex(key, field)
			fieldName := key
			dataType := field.Type().String()
			packagePath := field.Type().PkgPath()
			kind := field.Kind()

			fmt.Printf("%s %s Name: %s  Type: %s  Package: %s\n", logPrefix+" - ", kind, fieldName, dataType, packagePath)
			DisectFields(field.Type(), field, logPrefix+" -  - ", handlers.NewTagHandler(""))
		}

	case reflect.Slice:
		zeVal.Set(reflect.MakeSlice(zeVal.Type(), 1, 1))
		field := zeVal.Index(0).Addr()
		DisectFields(field.Type(), field, logPrefix+" - ", handlers.NewTagHandler(""))
		// fmt.Println("SLICEEEE", zeVal.Type().String(), reflect.SliceOf(zeVal.Type().Elem()), zeVal.Len())

	case reflect.Array:
		fmt.Printf("%+v\n", zeVal)
		for i := 0; i < zeVal.Len(); i++ {
			field := zeVal.Index(i).Addr()
			DisectFields(field.Type(), field, logPrefix+" - ", handlers.NewTagHandler(""))
		}

	default:
		// if zeVal.Type().String() == zeVal.Kind().String() {
		if fakeMaker.CanHandle(zeVal) {
			fakeMaker.GenerateValue(zeVal, tagHandler)
		} else {
			fmt.Println("BadSet", zeVal.Type().String(), zeVal.Kind().String())
		}
	}
}
