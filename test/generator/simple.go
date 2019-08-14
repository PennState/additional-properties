package acceptance

// Simple is the most basic struct to which we can add additional-
// properties.
type Simple struct {
	A  string
	AP map[string]interface{} `json:"*"`
}

func newTestSimple() Simple {
	return Simple{
		A: "Field A",
		AP: map[string]interface{}{
			"b": "Field B",
			"c": "Field C",
		},
	}
}
