package acceptance

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeneratedMarshalerWorks(t *testing.T) {
	expected, err := ioutil.ReadFile("simple.json")
	require.NoError(t, err)
	input := newTestSimple()
	actual, err := json.Marshal(&input)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestGeneratedUnmarshalerWorks(t *testing.T) {
	files := []string{"simple.json", "capitalized.json"}
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		require.NoError(t, err)
		var s Simple
		err = json.Unmarshal(data, &s)
		assert.NoError(t, err)
		assert.Equal(t, newTestSimple(), s)
	}
}
