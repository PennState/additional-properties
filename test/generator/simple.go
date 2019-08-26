package acceptance

// Simple is the most basic struct to which we can add additional-
// properties.
type Simple struct {
	FieldA string                 `json:"fieldA"`
	AP     map[string]interface{} `json:"*"`
}

func newZeroSimple() interface{} {
	return &Simple{}
}

func newTestSimple() interface{} {
	return &Simple{
		FieldA: "Field A",
		AP: map[string]interface{}{
			"fieldB": "Field B",
			"fieldC": "Field C",
		},
	}
}

func newTestSimpleWithoutAP() interface{} {
	return &Simple{
		FieldA: "Field A",
	}
}
