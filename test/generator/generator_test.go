package acceptance

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	Name          string
	MarshalJson   string
	UnmarshalJson string
	Data          func() interface{}
	Zero          func() interface{}
}{
	{Name: "Exact match", MarshalJson: "simple.json", UnmarshalJson: "simple.json", Data: newTestSimple, Zero: newZeroSimple},
	{Name: "Case insensitive", MarshalJson: "simple.json", UnmarshalJson: "capitalized.json", Data: newTestSimple, Zero: newZeroSimple},
	{Name: "No additional properties", MarshalJson: "noap.json", UnmarshalJson: "noap.json", Data: newTestSimpleWithoutAP, Zero: newZeroSimple},
	{Name: "Respects omitempty", MarshalJson: "omitempty.json", UnmarshalJson: "omitempty.json", Data: newTestOmitEmpty, Zero: newZeroOmitEmpty},
}

func TestGeneratedMarshalerWorks(t *testing.T) {
	for _, test := range tests {
		t.Log("Test name (marshaling): ", test.Name)
		expected, err := ioutil.ReadFile("./testdata/" + test.MarshalJson)
		require.NoError(t, err)
		input := test.Data()
		actual, err := json.Marshal(input)
		assert.NoError(t, err)
		assert.JSONEq(t, string(expected), string(actual))
	}
}

func TestGeneratedUnmarshalerWorks(t *testing.T) {
	for _, test := range tests {
		t.Log("Test name (unmarshaling): ", test.Name)
		data, err := ioutil.ReadFile("./testdata/" + test.UnmarshalJson)
		require.NoError(t, err)
		z := test.Zero()
		err = json.Unmarshal(data, z)
		assert.NoError(t, err)
		assert.EqualValues(t, test.Data(), z)
	}
}
