package handlers

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/fatih/structtag"
)

func TestEmptyTag(t *testing.T) {
	emptyTag := Tag{}
	if !emptyTag.IsEmpty() {
		t.Error("Tag should be empty")
	}

	tags, err := structtag.Parse(`derp:"thing,arg1,arg2"`)
	if err != nil {
		panic(err)
	}

	tag, _ := tags.Get("derp")

	nonEmptyTag := Tag{
		Tag: tag,
	}
	if nonEmptyTag.IsEmpty() {
		t.Error("Tag should not be empty")
	}
}

func TestFakerOverride(t *testing.T) {
	emptyTag := Tag{}
	if emptyTag.IsFakerOveridden() {
		t.Error("Faker should not be overriden")
	}

	tags, _ := structtag.Parse(`derp:"paragraph"`)
	tag, _ := tags.Get("derp")

	overiddenFakerTag := Tag{
		Tag: tag,
	}
	if !overiddenFakerTag.IsFakerOveridden() {
		t.Error("Faker should be overriden")
	}

	tags, _ = structtag.Parse(`derp:"thingdontexist"`)
	tag, _ = tags.Get("derp")

	overiddenFakerTag = Tag{
		Tag: tag,
	}
	if overiddenFakerTag.IsFakerOveridden() {
		t.Error("Faker should not be overriden")
	}
}

func TestTagOptions(t *testing.T) {
	arg1 := "arg1"
	tag := NewTagHandler(FakeItTag + `:"thing,` + arg1 + `,arg2"`)

	if !tag.HasIndex(0) || !tag.HasIndex(1) {
		t.Error("Option index should exist")
	}

	if tag.HasIndex(2) {
		t.Error("Option index should not exist")
	}

	if tag.GetIndex(0) != arg1 {
		t.Error("Option index value is not correct")
	}
}

func TestNewTagHandler(t *testing.T) {
	value1 := "value1"
	value2 := "value2"
	tag := NewTagHandler(FakeItTag + `:"thing,arg1=` + value1 + `,arg2=` + value2 + `"`)

	expectedOptions := []string{value1, value2}
	if tag.Options[0] != expectedOptions[0] || tag.Options[1] != expectedOptions[1] || len(tag.Options) != len(expectedOptions) {
		t.Error("Options were not initialised")
	}

	badTag := NewTagHandler(`:"=`)
	if !badTag.IsEmpty() || badTag.HasIndex(0) {
		t.Error("Tag should not have options or a tag")
	}

	badTag = NewTagHandler(`validTag:"argument"`)
	if !badTag.IsEmpty() || badTag.HasIndex(0) {
		t.Error("Tag should not have options or a tag")
	}
}

func TestFakerValue(t *testing.T) {
	digitLength := 7
	tag := NewTagHandler(FakeItTag + `:"digitsN,num=` + strconv.Itoa(digitLength) + `"`)
	value := tag.GetFakerValue()

	if value.Kind() != reflect.String || len(value.Interface().(string)) != digitLength {
		t.Error("Faker value should have been returned")
	}

}

func TestTypeHandlerFakeIt(t *testing.T) {
	expectedString := "schleeble"
	handler := TypeHandler{
		Kind: reflect.String,
		Type: "string",
	}
	handler.GetDefaultFaker = func() reflect.Value {
		return reflect.ValueOf(expectedString)
	}
	dummyString := "string"
	value := reflect.ValueOf(&dummyString).Elem()
	handler.FakeIt(value, NewTagHandler(""))

	if value.Interface().(string) != expectedString {
		t.Error("Default handler should have been used")
	}

	handler.FakeIt(value, NewTagHandler(FakeItTag+`:"paragraph"`))
	if value.Interface().(string) == expectedString {
		t.Error("Default handler should not have been used")
	}

}
