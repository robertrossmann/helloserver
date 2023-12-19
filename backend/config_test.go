package backend

import (
	"helloserver/testtools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	testtools.Bootstrap()
}

func Test_NewConfig(t *testing.T) {
	cfg, err := NewConfig()

	assert.NoError(t, err)
	assert.Equal(t, cfg.Port, 3000)
}
