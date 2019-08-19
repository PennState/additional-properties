package aputil

import (
	"go/ast"
	"reflect"
	"strings"
)

const tagName = "json"

type Tag struct {
	Name    string
	Options []string
}

func NewTagFromField(f ast.Field) (Tag, bool) {
	for _, txt := range strings.Split(f.Tag.Value, " ") {
		if strings.HasPrefix(txt, "`"+tagName) {
			// TODO: this is kludgey
			return NewTag(strings.TrimSuffix(strings.TrimPrefix(txt, "`"+tagName+":\""), "\"`"))
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
	tkns := strings.Split(txt, ",")
	if len(tkns) == 1 {
		return Tag{
			Name: tkns[0],
		}, true
	}
	return Tag{
		Name:    tkns[0],
		Options: tkns[1:],
	}, true
}

func GetJSONName(f reflect.StructField) string {
	tag, ok := NewTagFromStructField(f)
	if ok {
		return tag.Name
	}
	return f.Name
}
