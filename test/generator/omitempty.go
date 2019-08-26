package acceptance

// Simple is the most basic struct to which we can add additional-
// properties.
type OmitEmpty struct {
	FieldA string                 `json:"fieldA,omitempty"`
	AP     map[string]interface{} `json:"*"`
}

func newTestOmitEmpty() Simple {
	return Simple{
		FieldA: "",
		AP: map[string]interface{}{
			"fieldB": "Field B",
			"fieldC": "Field C",
		},
	}
}
