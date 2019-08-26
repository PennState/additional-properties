package aputil

import (
	"go/ast"
	"reflect"
	"strings"
)

const tagName = "json"

type Tag struct {
	Name    string
	Options map[string]bool
}

func NewTagFromField(f *ast.Field) (Tag, bool) {
	if f.Tag == nil {
		return Tag{}, false
	}
	tkns := strings.Split(strings.TrimSuffix(strings.TrimPrefix(f.Tag.Value, "`"), "`"), " ")
	for _, tkn := range tkns {
		if strings.HasPrefix(tkn, tagName) {
			return NewTag(strings.TrimPrefix(tkn, tagName+":"))
		}
	}
	return Tag{}, false
}

func NewTagFromStructField(f reflect.StructField) (Tag, bool) {
	txt, ok := f.Tag.Lookup(tagName)
	if !ok {
		return Tag{}, ok
	}
	return NewTag(txt)
}

func NewTag(txt string) (Tag, bool) {
	tkns := strings.Split(strings.TrimSuffix(strings.TrimPrefix(txt, "\""), "\""), ",")
	opts := map[string]bool{}
	for _, opt := range tkns[1:] {
		opts[opt] = true
	}
	return Tag{
		Name:    tkns[0],
		Options: opts,
	}, true
}

func GetJSONName1(f reflect.StructField) string {
	tag, ok := NewTagFromStructField(f)
	if ok {
		return tag.Name
	}
	return f.Name
}

func GetJSONName(f *ast.Field) string {
	tag, ok := NewTagFromField(f)
	if ok {
		return tag.Name
	}
	return f.Names[0].Name
}
