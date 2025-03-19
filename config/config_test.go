package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	configJSON := `{"port": "8080", "log_level": "Info"}`
	err := os.WriteFile("test_config.json", []byte(configJSON), 0644)
	assert.NoError(t, err)
	defer os.Remove("test_config.json")

	config, err := LoadConfig("test_config.json")
	assert.NoError(t, err)
	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, "Info", config.LogLevel)
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	_, err := LoadConfig("non_existent.json")
	assert.Error(t, err)
}
