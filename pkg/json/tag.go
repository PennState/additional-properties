package json

import (
	"reflect"
	"strings"
)

const (
	tagId string = "json"
)

type Tag struct {
	Name      string
	Omit      bool
	OmitEmpty bool
	String    bool
	Wildcard  bool
}

func NewTag(sf reflect.StructField) Tag {
	ts := sf.Tag.Get(tagId)

	if ts == "-" {
		return Tag{
			Omit: true,
		}
	}

	if ts == "*" {
		return Tag{
			Omit:     true,
			Wildcard: true,
		}
	}

	idx := strings.Index(ts, ",")
	if idx == -1 {
		return Tag{
			Name: ts,
		}
	}

	c := make(map[string]bool)
	for _, k := range strings.Split(ts, ",") {
		c[k] = true
	}

	return Tag{
		Name:      ts[:idx],
		OmitEmpty: c["omitempty"],
		String:    c["string"],
	}
}
