package acceptance

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/PennState/additional-properties/pkg/generator"
	"github.com/PennState/proctor/pkg/goldenfile"
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
		input := test.Data()
		actual, err := json.Marshal(input)
		assert.NoError(t, err)
		fp := goldenfile.GetDefaultFilePath(test.MarshalJson)
		goldenfile.AssertJSONEq(t, fp, string(actual))
	}
}

func TestGeneratedUnmarshalerWorks(t *testing.T) {
	for _, test := range tests {
		t.Log("Test name (unmarshaling): ", test.Name)
		fp := goldenfile.GetDefaultFilePath(test.UnmarshalJson)
		data, err := ioutil.ReadFile(fp)
		require.NoError(t, err)
		z := test.Zero()
		err = json.Unmarshal(data, z)
		assert.NoError(t, err)
		assert.EqualValues(t, test.Data(), z)
	}
}

func TestGeneratedFileMatches(t *testing.T) {
	tests := []struct {
		Name string
		File string
	}{
		{Name: "Omit Empty Generated File", File: "omitempty_gen.go"},
		{Name: "Simple JSON File", File: "simple_gen.go"},
	}
	t.Log(os.Args)
	err := generator.Run()
	require.NoError(t, err)
	for _, test := range tests {
		actual, err := ioutil.ReadFile(test.File)
		require.NoError(t, err)
		fp := goldenfile.GetDefaultFilePath(test.File)
		goldenfile.AssertStringEq(t, fp, string(actual))
	}
}
