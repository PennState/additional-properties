package ap_test

import "encoding/json"

func NewZeroAnonymous() interface{} {
	return &struct {
		*Alias
	}{
		Alias: (*Alias)(&Real{}),
	}
}

func NewTestAnonymous() interface{} {
	return struct {
		Alias
	}{
		Alias: (Alias)(Real{
			FieldA: "Field A",
			AP: map[string]json.RawMessage{
				"fieldB": json.RawMessage([]byte("\"Field B\"")),
				"fieldC": json.RawMessage([]byte("\"Field C\"")),
			}}),
	}
}
