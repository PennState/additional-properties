package ap_test

import "encoding/json"

type Real struct {
	FieldA string                     `json:"fieldA"`
	AP     map[string]json.RawMessage `json:"*"`
}

type Alias Real

// alias := struct {
// 	*Alias
// }{
// 	Alias: (*Alias)(u),
// }

func NewZeroAlias() interface{} {
	z := (Alias)(Real{})
	return &z
}

func NewTestAlias() interface{} {
	t := (Alias)(Real{
		FieldA: "Field A",
		AP: map[string]json.RawMessage{
			"fieldB": json.RawMessage([]byte("\"Field B\"")),
			"fieldC": json.RawMessage([]byte("\"Field C\"")),
		},
	})
	return &t
}
