package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCloseEnoughValue(t *testing.T) {
	input := []int{0, 10, 20, 30, 40, 100, 1000000}

	value, index := FindCloseEnoughValue(input, 21, 2.0)
	assert.Equal(t, 20, value)
	assert.Equal(t, 2, index)

	value, index = FindCloseEnoughValue(input, 100, 10.0)
	assert.Equal(t, 100, value)
	assert.Equal(t, 5, index)

	value, index = FindCloseEnoughValue(input, 1000, 5.0)
	assert.Equal(t, 0, value)
	assert.Equal(t, 0, index)
}
