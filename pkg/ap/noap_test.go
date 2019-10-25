package ap_test

type NoAP struct {
	FieldA string `json:"fieldA"`
}

func NewZeroNoAP() interface{} {
	return &NoAP{}
}

func NewTestNoAP() interface{} {
	return &NoAP{
		FieldA: "Field A",
	}
}
