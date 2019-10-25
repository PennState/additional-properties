package ap_test

import "encoding/json"

// Simple is the most basic struct to which we can add additional-
// properties.
type Simple struct {
	FieldA string                     `json:"fieldA"`
	AP     map[string]json.RawMessage `json:"*"`
}

func NewZeroSimple() interface{} {
	return &Simple{}
}

func NewTestSimple() interface{} {
	return &Simple{
		FieldA: "Field A",
		AP: map[string]json.RawMessage{
			"fieldB": json.RawMessage([]byte("\"Field B\"")),
			"fieldC": json.RawMessage([]byte("\"Field C\"")),
		},
	}
}

func NewTestSimpleWithoutAP() interface{} {
	return &Simple{
		FieldA: "Field A",
		AP:     map[string]json.RawMessage{},
	}
}
