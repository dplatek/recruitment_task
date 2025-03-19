package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	configJSON := `{"port": "8080", "log_level": "Info"}`
	err := os.WriteFile("testconfig.json", []byte(configJSON), 0644)
	assert.NoError(t, err)
	defer os.Remove("testconfig.json")

	config, err := LoadConfig("testconfig.json")
	assert.NoError(t, err)
	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, "Info", config.LogLevel)
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	_, err := LoadConfig("non_existent.json")
	assert.Error(t, err)
}
