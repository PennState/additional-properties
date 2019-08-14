package acceptance

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	data, err := ioutil.ReadFile("simple.json")
	require.NoError(t, err)
	var s Simple
	err = json.Unmarshal(data, &s)
	assert.NoError(t, err)
	log.Info("Simple: ", s)
	assert.True(t, false)
}
