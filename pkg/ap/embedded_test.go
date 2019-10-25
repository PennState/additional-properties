package ap_test

import "encoding/json"

// Inner is the most basic struct to which we can add additional-
// properties.
type Inner struct {
	FieldA string                     `json:"fieldA"`
	AP     map[string]json.RawMessage `json:"*"`
}

type Outer struct {
	Inner
	FieldD string `json:"fieldD"`
}

func NewZeroOuter() interface{} {
	return &Outer{
		Inner: Inner{},
	}
}

func NewTestOuter() interface{} {
	return &Outer{
		Inner: Inner{
			FieldA: "Field A",
			AP: map[string]json.RawMessage{
				"fieldB": json.RawMessage([]byte("\"Field B\"")),
				"fieldC": json.RawMessage([]byte("\"Field C\"")),
			},
		},
		FieldD: "Field D",
	}
}
