package json

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestArraysAreDereferenced(t *testing.T) {
	assert := assert.New(t)

	a := []string{"Test"}
	p := &a
	pp := &p
	ppp := &pp

	assert.Equal(reflect.String, dereferencedKind(a))
	assert.Equal(reflect.String, dereferencedKind(p))
	assert.Equal(reflect.String, dereferencedKind(pp))
	assert.Equal(reflect.String, dereferencedKind(ppp))
}

func TestPointersToPointersAreDereferenced(t *testing.T) {
	assert := assert.New(t)

	s := "Test"
	p := &s
	pp := &p
	ppp := &pp

	assert.Equal(reflect.String, dereferencedKind(p))
	assert.Equal(reflect.String, dereferencedKind(pp))
	assert.Equal(reflect.String, dereferencedKind(ppp))
}

func TestAdditionalPropertiesFieldOnNonStruct(t *testing.T) {
	v := "Test"
	testFunc := func() {
		_, _ = additionalPropertiesField(v)
	}
	assert.Panics(t, testFunc)
}

func TestAdditionalPropertiesFieldNotDefinedInStruct(t *testing.T) {
	type testStruct struct {
		S string
	}
	var ts testStruct
	ap, err := additionalPropertiesField(ts)
	assert.NoError(t, err)
	assert.Empty(t, ap)
	ap["miscellaneous"] = json.RawMessage("\"with arbitrary content\"")
	assert.Len(t, ap, 1)
}

func TestAdditionalPropertiesFieldDefinedInStruct(t *testing.T) {
	type testStruct struct {
		S  string
		AP map[string]json.RawMessage `json:"*"`
	}
	var ts testStruct
	ts.AP = make(map[string]json.RawMessage)
	ts.AP["miscellaneous"] = json.RawMessage("\"with arbitrary content\"")
	ap, err := additionalPropertiesField(ts)
	assert.NoError(t, err)
	assert.Len(t, ap, 1)
	ap["second value"] = json.RawMessage("\"more arbitrary content\"")
	assert.Len(t, ap, 2)
}

func TestAdditionalPropertiesFieldDefinedButUnexportedInStruct(t *testing.T) {
	type testStruct struct {
		S  string
		ap map[string]json.RawMessage `json:"*"`
	}
	var ts testStruct
	ts.ap = make(map[string]json.RawMessage)
	ts.ap["miscellaneous"] = json.RawMessage("\"with arbitrary content\"")
	ap, err := additionalPropertiesField(ts)
	assert.Error(t, err)
	assert.Nil(t, ap)
}

func TestAdditionalPropertiesFieldNotMapStringJsonRawMessage(t *testing.T) {
	type testStruct struct {
		S string `json:"*"`
	}
	var ts testStruct
	ap, err := additionalPropertiesField(ts)
	assert.Error(t, err)
	assert.Nil(t, ap)
}

func TestJsonNameOnTaggedField(t *testing.T) {
	type testStruct struct {
		F string `json:"f"`
	}
	var ts testStruct
	sf, _ := reflect.TypeOf(ts).FieldByName("F")
	n, ok := jsonName(sf)
	assert.True(t, ok)
	assert.Equal(t, "f", n)
}

func TestJsonNameOnTaggedFieldWithOtherParts(t *testing.T) {
	type testStruct struct {
		F string `json:"f,omitempty"`
	}
	var ts testStruct
	sf, _ := reflect.TypeOf(ts).FieldByName("F")
	n, ok := jsonName(sf)
	assert.True(t, ok)
	assert.Equal(t, "f", n)
}

func TestJsonNameOnTaggedFieldWithMinusFieldName(t *testing.T) {
	type testStruct struct {
		F string `json:"-,"`
	}
	var ts testStruct
	sf, _ := reflect.TypeOf(ts).FieldByName("F")
	n, ok := jsonName(sf)
	assert.True(t, ok)
	assert.Equal(t, "-", n)
}

func TestJsonNameOnTaggedFieldWithAsteriskFieldName(t *testing.T) {
	type testStruct struct {
		F string `json:"*,"`
	}
	var ts testStruct
	sf, _ := reflect.TypeOf(ts).FieldByName("F")
	n, ok := jsonName(sf)
	assert.True(t, ok)
	assert.Equal(t, "*", n)
}

func TestJsonNameOnUntaggedField(t *testing.T) {
	type testStruct struct {
		F string
	}
	var ts testStruct
	sf, _ := reflect.TypeOf(ts).FieldByName("F")
	n, ok := jsonName(sf)
	assert.True(t, ok)
	assert.Equal(t, "F", n)
}

func TestJsonNameOnSkippedField(t *testing.T) {
	type testStruct struct {
		F string `json:"-"`
	}
	var ts testStruct
	sf, _ := reflect.TypeOf(ts).FieldByName("F")
	n, ok := jsonName(sf)
	assert.False(t, ok)
	assert.Empty(t, n)
}

func TestJsonNameOnWildcardField(t *testing.T) {
	type testStruct struct {
		F string `json:"*"`
	}
	var ts testStruct
	sf, _ := reflect.TypeOf(ts).FieldByName("F")
	n, ok := jsonName(sf)
	assert.False(t, ok)
	assert.Empty(t, n)
}
