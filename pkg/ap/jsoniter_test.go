package ap_test

import (
	"io/ioutil"
	"testing"

	"github.com/PennState/additional-properties/pkg/ap"
	"github.com/PennState/proctor/pkg/goldenfile"
	_ "github.com/PennState/proctor/pkg/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:gochecknoglobals
var cases = []struct {
	Name    string
	JSONOut string
	JSONIn  string
	Data    func() interface{}
	Zero    func() interface{}
}{
	{"Exact match", "simple.json", "simple.json", NewTestSimple, NewZeroSimple},
	{"Case insensitive", "simple.json", "capitalized.json", NewTestSimple, NewZeroSimple},
	{"No AP field", "noap.json", "noap.json", NewTestNoAP, NewZeroNoAP},
	{"No additional properties", "noap.json", "noap.json", NewTestSimpleWithoutAP, NewZeroSimple},
	{"Respects omitempty", "omitempty.json", "omitempty.json", NewTestOmitEmpty, NewZeroOmitEmpty},
	{"Embedded struct with AP", "embedded.json", "embedded.json", NewTestOuter, NewZeroOuter},
	{"Aliased AP", "aliased.json", "aliased.json", NewTestAlias, NewZeroAlias},
	// {"Anonymous aliased AP", "anonymous.json", "anonymous.json", NewTestAnonymous, NewZeroAnonymous},
}

func TestMarshaling(t *testing.T) {
	json := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	}.Froze()
	ap.RegisterAdditionalPropertiesExtension(json)
	for idx := range cases {
		c := cases[idx]
		t.Run(c.Name, func(t *testing.T) {
			input := c.Data()
			actual, err := json.Marshal(input)
			assert.NoError(t, err)
			fp := goldenfile.GetDefaultFilePath(c.JSONOut)
			goldenfile.AssertJSONEq(t, fp, string(actual))
		})
	}
}

func TestUnmarshaling(t *testing.T) {
	json := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	}.Froze()
	ap.RegisterAdditionalPropertiesExtension(json)
	for idx := range cases {
		c := cases[idx]
		t.Run(c.Name, func(t *testing.T) {
			fp := goldenfile.GetDefaultFilePath(c.JSONIn)
			data, err := ioutil.ReadFile(fp)
			require.NoError(t, err)
			z := c.Zero()
			err = json.Unmarshal(data, z)
			assert.NoError(t, err)
			assert.EqualValues(t, c.Data(), z)
		})
	}
}
