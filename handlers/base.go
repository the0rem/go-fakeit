package handlers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structtag"
	"github.com/the0rem/go-fakeit/fakers"
)

const (
	// FakeItTag identifies the tag uses on Struct Fields
	// for specifying which fakeit method to use
	FakeItTag = "fakeit"
)

var fakerHandlers = fakers.NewFakers()

// NewTagHandler takes a given Struct Field Tag and populates a Tag struct with the processed tag data.
func NewTagHandler(tagString string) *Tag {
	var tag *structtag.Tag
	tags, err := structtag.Parse(string(tagString))
	if err != nil {
		panic(err)
	}

	tag, err = tags.Get(FakeItTag)
	if err != nil {
		tag = &structtag.Tag{}
	} else {
		fmt.Println(tag)         // Output: json:"foo,omitempty,string"
		fmt.Println(tag.Key)     // Output: json
		fmt.Println(tag.Name)    // Output: foo
		fmt.Println(tag.Options) // Output: [omitempty string]
	}

	var options []string

	for _, option := range tag.Options {
		options = append(options, parseOption(option))
	}

	fmt.Println(options)

	return &Tag{
		Options: options,
		Tag:     tag,
	}
}

func parseOption(option string) string {
	optionValues := strings.Split(option, "=")
	if len(optionValues) > 1 {
		return strings.Join(optionValues[1:], "")
	}
	return optionValues[0]
}

type TypeHandler struct {
	Kind            reflect.Kind
	Type            string
	Package         string
	GetDefaultFaker func() reflect.Value
}

func (typeHandler *TypeHandler) FakeIt(field reflect.Value, tag *Tag) {
	if tag.IsFakerOveridden() {
		fmt.Println("Shit is overridden")
		field.Set(tag.GetFakerValue())
	} else {
		field.Set(typeHandler.GetDefaultFaker())
	}
}

type Tag struct {
	Options []string
	Tag     *structtag.Tag
}

func (tag *Tag) IsEmpty() bool {
	return tag.Tag.Key != ""
}

func (tag *Tag) HasIndex(index int) bool {
	return len(tag.Options) >= index
}

func (tag *Tag) GetIndex(index int) string {
	return tag.Options[index]
}

func (tag *Tag) IsFakerOveridden() bool {
	if _, ok := fakerHandlers[tag.Tag.Name]; ok {
		return true
	}
	return false
}

func (tag *Tag) GetFakerValue() reflect.Value {
	faker := fakerHandlers[tag.Tag.Name]
	fakerValue := reflect.ValueOf(faker)
	var options []reflect.Value
	for _, option := range tag.Options {
		options = append(options, reflect.ValueOf(option))
	}
	functionResponse := fakerValue.Call(options)
	return reflect.ValueOf(functionResponse[0].Interface())
}
