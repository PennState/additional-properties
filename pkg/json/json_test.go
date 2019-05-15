package json

import (
	std_json "encoding/json"
	"reflect"
	"testing"

	//	"github.com/PennState/go-additional-properties/pkg/json"

	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

//
//Test Marshaling
//

func TestEmbeddedStructFieldsAreAddedToParent(t *testing.T) {
	type A struct {
		A string
	}
	type B struct {
		A
		B string
	}
	type C struct {
		B
		C string
	}

	v := C{
		B: B{
			A: A{
				A: "A string",
			},
			B: "B string",
		},
		C: "C string",
	}

	json, err := Marshal(v)
	if err != nil {
		log.Error(err)
	}
	log.Info(string(json))
	assert.Equal(t, "{\"A\":\"A string\",\"B\":\"B string\",\"C\":\"C string\"}", string(json))
}

//
//Test Unmarshaling
//

//
//Test DereferenceKind method
//

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

//
//Test AdditionalPropertiesField method
//

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
	ap["miscellaneous"] = std_json.RawMessage("\"with arbitrary content\"")
	assert.Len(t, ap, 1)
}

func TestAdditionalPropertiesFieldDefinedInStruct(t *testing.T) {
	type testStruct struct {
		S  string
		AP map[string]std_json.RawMessage `json:"*"`
	}
	var ts testStruct
	ts.AP = make(map[string]std_json.RawMessage)
	ts.AP["miscellaneous"] = std_json.RawMessage("\"with arbitrary content\"")
	ap, err := additionalPropertiesField(ts)
	assert.NoError(t, err)
	assert.Len(t, ap, 1)
	ap["second value"] = std_json.RawMessage("\"more arbitrary content\"")
	assert.Len(t, ap, 2)
}

func TestAdditionalPropertiesFieldDefinedButUnexportedInStruct(t *testing.T) {
	type testStruct struct {
		S  string
		ap map[string]std_json.RawMessage `json:"*"`
	}
	var ts testStruct
	ts.ap = make(map[string]std_json.RawMessage)
	ts.ap["miscellaneous"] = std_json.RawMessage("\"with arbitrary content\"")
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

//
//Test isEmpty method
//

func TestIsEmptyWithNestedStruct(t *testing.T) {
	type AStruct struct {
		AField  string
		BStruct struct {
			BField int
		}
	}
	var s AStruct
	p := &s
	pp := &p
	assert.True(t, isEmpty(reflect.ValueOf(s)))
	assert.True(t, isEmpty(reflect.ValueOf(p)))
	assert.True(t, isEmpty(reflect.ValueOf(pp)))
}

//
//Test jsonName method
//

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

func TestMarshalStruct(t *testing.T) {
	type testStructA struct {
		A string
		B string `json:"b"`
		C string `json:"-"`
		d string
		E struct {
			EA string
			EB string `json:"eb"`
			EC string `json:"-"`
		}
		F []struct {
			FA string `json:"fa"`
		}
		G map[string]std_json.RawMessage `json:"*"`
	}
}

func TestMarshalStructSkipsUnexportedFields(t *testing.T) {
	testStruct := struct {
		A string
		b string
		C struct {
			D string
			e string
		}
	}{
		A: "A",
		b: "b",
		C: struct {
			D string
			e string
		}{
			D: "D",
			e: "e",
		},
	}
	data, err := marshalStruct(testStruct)
	assert.NoError(t, err)
	assert.Equal(t, "{\"A\":\"A\",\"C\":{\"D\":\"D\"}}", string(data))
}
