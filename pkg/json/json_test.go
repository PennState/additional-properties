package json

import (
	"encoding/json"
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
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
	testFunc := func () {
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
		S string
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
		S string
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

