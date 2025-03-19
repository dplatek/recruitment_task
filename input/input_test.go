package input

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInputFromFile(t *testing.T) {
	inputContent := "10\n20\n30\n40\n"
	err := os.WriteFile("testinput.txt", []byte(inputContent), 0644)
	assert.NoError(t, err)
	defer os.Remove("testinput.txt")

	input, err := LoadInputFromFile("testinput.txt")
	assert.NoError(t, err)
	assert.Equal(t, []int{10, 20, 30, 40}, input)
}

func TestLoadInputFromFile_InvalidInput(t *testing.T) {
	inputContent := "10\ninvalid\n30\n"
	err := os.WriteFile("testinvalid.txt", []byte(inputContent), 0644)
	assert.NoError(t, err)
	defer os.Remove("testinvalid.txt")

	_, err = LoadInputFromFile("testinvalid.txt")
	assert.Error(t, err)
}

func TestLoadInputFromFile_FileNotFound(t *testing.T) {
	_, err := LoadInputFromFile("nonexistent.txt")
	assert.Error(t, err)
}
