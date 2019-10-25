package ap_test

import (
	"encoding/json"
	"time"
)

// OmitEmpty is the most basic struct to which we can add additional-
// properties.
type OmitEmpty struct {
	A string  `json:"fieldA,omitempty"`
	B int64   `json:"fieldB,omitempty"`
	C float64 `json:"fieldC,omitempty"`
	//	D  complex128                 `json:"fieldD,omitempty"`
	E  bool                       `json:"fieldE,omitempty"`
	F  map[string]string          `json:"fieldF,omitempty"`
	G  *map[string]string         `json:"fieldG,omitempty"`
	H  []string                   `json:"fieldH,omitempty"`
	I  *[]string                  `json:"fieldI,omitempty"`
	J  EmptyStruct                `json:"fieldJ,omitempty"`
	K  *EmptyStruct               `json:"fieldK,omitempty"`
	L  time.Time                  `json:"fieldL,omitempty"`
	M  string                     `json:"fieldM"`
	Z  string                     `json:"fieldZ"`
	AP map[string]json.RawMessage `json:"*"`
}

type EmptyStruct struct {
	Y string `json:"fieldY,omitempty"`
}

func NewZeroOmitEmpty() interface{} {
	return &OmitEmpty{}
}

func NewTestOmitEmpty() interface{} {
	return &OmitEmpty{
		Z: "Field Z",
		AP: map[string]json.RawMessage{
			"ap1": json.RawMessage([]byte("\"AP 1\"")),
			"ap2": json.RawMessage([]byte("\"AP 2\"")),
		},
	}
}
