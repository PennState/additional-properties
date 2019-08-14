package acceptance

type Simple struct {
	A  string
	AP map[string]interface{} `json:"*"`
}
