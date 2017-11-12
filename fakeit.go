package fakeit

import (
	"fmt"
	"reflect"

	"github.com/the0rem/go-fakeit/fakers"
	"github.com/the0rem/go-fakeit/handlers"
)

var typeHandlers = []*handlers.TypeHandler{
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
}

// FakeIt takes a struct and populates it with fake data.
func FakeIt(a interface{}) {
	fakeMaker := NewFakeMaker()
	fakeMaker.FakeIt(a)
}

// NewFakeMaker return a new FakeMaker struct.
func NewFakeMaker() *FakeMaker {
	return &FakeMaker{
		Fakers:                 fakers.NewFakers(),
		MaxRecursion:           -1,
		AllowCircularRecursion: false,
		TypeHandlers:           typeHandlers,
		Verbose:                false,
	}
}

// FakeMaker handles filling variables with fake data
type FakeMaker struct {
	Fakers                 map[string]func(args ...interface{}) interface{}
	TypeHandlers           []*handlers.TypeHandler
	MaxRecursion           int
	recusrionDepth         int
	followedStructs        []string
	AllowCircularRecursion bool
	Verbose                bool
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

	if !field.CanAddr() || !field.CanSet() || tag.IsFieldIgnored() {
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

func (fakeMaker *FakeMaker) FakeIt(a interface{}) {
	t := reflect.TypeOf(a)
	valueOf := reflect.ValueOf(a)
	followedStructs := []string{}
	fakeMaker.fakeIt(t, valueOf, handlers.NewTagHandler(""), followedStructs)
}

// fakeIt takes the reflect properties of an interface and recursively populates the
// associated proerties using the appropriate fakers.
func (fakeMaker *FakeMaker) fakeIt(t reflect.Type, valueOf reflect.Value, tagHandler *handlers.Tag, followedStructs []string) {

	if valueOf.Kind() != reflect.Ptr {
		panic("Aint a pointer: " + valueOf.Kind().String())
	}

	if fakeMaker.recusrionDepth == fakeMaker.MaxRecursion {
		return
	}

	zeVal := reflect.Indirect(valueOf)
	kind := zeVal.Kind()

	if fakeMaker.CanHandle(zeVal) {
		fakeMaker.GenerateValue(zeVal, tagHandler)
		// return
	}

	if kind == reflect.Struct || kind == reflect.Map || kind == reflect.Interface || kind == reflect.Slice || kind == reflect.Array {
		fakeMaker.recusrionDepth++
	}

	switch kind {

	case reflect.Ptr:
		if !zeVal.IsValid() {
			fmt.Println("Aint valid")
			return
		}

		thing := reflect.New(zeVal.Type().Elem())

		if !zeVal.CanSet() {
			return
		}

		if !fakeMaker.AllowCircularRecursion && stringInSlice(thing.Elem().Type().String(), followedStructs) {
			return
		}

		zeVal.Set(thing)
		fakeMaker.fakeIt(thing.Type(), thing, tagHandler, followedStructs)

	case reflect.Struct:

		// if fakeMaker.Verbose {
		fmt.Printf("%d %s Type: %s\n", fakeMaker.recusrionDepth, zeVal.Kind(), zeVal.Type())
		// }

		if !fakeMaker.AllowCircularRecursion && stringInSlice(zeVal.Type().String(), followedStructs) {
			fmt.Println("Had to stop going in circles", zeVal.Type().String(), followedStructs)
			return
		}

		followedStructs = append(followedStructs, zeVal.Type().String())

		for j := 0; j < zeVal.NumField(); j++ {

			field := zeVal.Field(j).Addr()
			fieldName := zeVal.Type().Field(j).Name
			tag := handlers.NewTagHandler(string(zeVal.Type().Field(j).Tag))

			dataType := field.Type().String()
			packagePath := field.Type().PkgPath()
			kind := field.Kind()

			// if fakeMaker.Verbose {
			fmt.Printf("%d %s Name: %s Type: %s Tag: %s Package: %s\n", fakeMaker.recusrionDepth, kind, fieldName, dataType, tag, packagePath)
			// }

			if tag.IsFieldIgnored() {
				fmt.Println("Im ignored")
				continue
			}

			fakeMaker.fakeIt(field.Type(), field, tag, followedStructs)
		}

	case reflect.Map:
		zeVal.Set(reflect.MakeMap(zeVal.Type()))
		keys := zeVal.MapKeys()
		if fakeMaker.Verbose {
			fmt.Println("MAPPPPPP", zeVal.Type().String(), reflect.SliceOf(zeVal.Type().Elem()), zeVal.Len(), keys)
		}

		for _, key := range keys {
			field := zeVal.MapIndex(key).Addr()
			// field := reflect.New(zeVal.Type()).Elem()
			fakeMaker.fakeIt(field.Type(), field, handlers.NewTagHandler(""), followedStructs)
			zeVal.SetMapIndex(key, field)
			fieldName := key
			dataType := field.Type().String()
			packagePath := field.Type().PkgPath()
			kind := field.Kind()

			if fakeMaker.Verbose {
				fmt.Printf("%d %s Name: %s  Type: %s  Package: %s\n", fakeMaker.recusrionDepth, kind, fieldName, dataType, packagePath)
			}
			fakeMaker.fakeIt(field.Type(), field, handlers.NewTagHandler(""), followedStructs)
		}

	case reflect.Slice:
		zeVal.Set(reflect.MakeSlice(zeVal.Type(), 1, 1))
		field := zeVal.Index(0).Addr()
		fakeMaker.fakeIt(field.Type(), field, handlers.NewTagHandler(""), followedStructs)

	case reflect.Array:
		for i := 0; i < zeVal.Len(); i++ {
			field := zeVal.Index(i).Addr()
			fakeMaker.fakeIt(field.Type(), field, handlers.NewTagHandler(""), followedStructs)
		}

	default:
		if fakeMaker.CanHandle(zeVal) {
			fakeMaker.GenerateValue(zeVal, tagHandler)
		} else {
			fmt.Println("BadSet", zeVal.Type().String(), zeVal.Kind().String())
		}
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
