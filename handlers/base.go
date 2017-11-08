package handlers

import "reflect"

type Liar struct {
	Kind    reflect.Kind
	Type    string
	Package string
	Fill    func(field reflect.Value, args Tag)
}

type Tag struct {
	Args []string
}

func (tag *Tag) HasIndex(index int) bool {
	return len(tag.Args) >= index
}

func (tag *Tag) GetIndex(index int) string {
	return tag.Args[index]
}
