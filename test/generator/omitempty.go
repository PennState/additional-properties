package acceptance

import "time"

// Simple is the most basic struct to which we can add additional-
// properties.
type OmitEmpty struct {
	A  string                 `json:"fieldA,omitempty"`
	B  int64                  `json:"fieldB,omitempty"`
	C  float64                `json:"fieldC,omitempty"`
	D  complex128             `json:"fieldD,omitempty"`
	E  bool                   `json:"fieldE,omitempty"`
	F  map[string]string      `json:"fieldF,omitempty"`
	G  *map[string]string     `json:"fieldG,omitempty"`
	H  []string               `json:"fieldH,omitempty"`
	I  *[]string              `json:"fieldI,omitempty"`
	J  EmptyStruct            `json:"fieldJ,omitempty"`
	K  *EmptyStruct           `json:"fieldK,omitempty"`
	L  time.Time              `json:"fieldL,omitempty"`
	Z  string                 `json:"fieldZ"`
	AP map[string]interface{} `json:"*"`
}

type EmptyStruct struct {
	Y string `json:"fieldZ,omitempty"`
}

func newZeroOmitEmpty() interface{} {
	return &OmitEmpty{}
}

func newTestOmitEmpty() interface{} {
	return &OmitEmpty{
		Z: "Field Z",
		AP: map[string]interface{}{
			"ap1": "AP 1",
			"ap2": "AP 2",
		},
	}
}
