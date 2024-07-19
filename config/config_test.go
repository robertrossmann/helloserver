package config

import (
	"helloserver/testtools"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	testtools.Bootstrap()
}

func Test_NewConfig(t *testing.T) {
	cfg, err := New()

	assert.NoError(t, err)
	assert.Equal(t, cfg.Port, 3000)
	assert.Equal(t, cfg.LogLevel, slog.LevelInfo)
}
