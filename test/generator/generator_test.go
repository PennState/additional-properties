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
	Data          func() Simple
}{
	{Name: "Exact match", MarshalJson: "simple.json", UnmarshalJson: "simple.json", Data: newTestSimple},
	{Name: "Case insensitive", MarshalJson: "simple.json", UnmarshalJson: "capitalized.json", Data: newTestSimple},
	{Name: "No additional properties", MarshalJson: "noap.json", UnmarshalJson: "noap.json", Data: newTestSimpleWithoutAP},
	{Name: "Respects omitempty", MarshalJson: "omitempty.json", UnmarshalJson: "omitempty.json", Data: newTestOmitEmpty},
}

func TestGeneratedMarshalerWorks(t *testing.T) {
	for _, test := range tests {
		t.Log("Test name (marshaling): ", test.Name)
		expected, err := ioutil.ReadFile("./testdata/" + test.MarshalJson)
		require.NoError(t, err)
		input := test.Data()
		actual, err := json.Marshal(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestGeneratedUnmarshalerWorks(t *testing.T) {
	for _, test := range tests {
		t.Log("Test name (unmarshaling): ", test.Name)
		data, err := ioutil.ReadFile("./testdata/" + test.UnmarshalJson)
		require.NoError(t, err)
		var s Simple
		err = json.Unmarshal(data, &s)
		assert.NoError(t, err)
		assert.Equal(t, test.Data(), s)
	}
}
