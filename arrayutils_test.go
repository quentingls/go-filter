package gofilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	input := []float64{0.0, 0.0, 1.0, 0.0, 0.0}
	expected := []float64{0.0, 0.0, 0.0, 1.0, 0.0}
	assert.Equal(t, expected, rotate(input, 1))

	expected = []float64{0.0, 0.0, 1.0, 0.0, 0.0}
	assert.Equal(t, expected, rotate(input, 5))

	expected = []float64{0.0, 1.0, 0.0, 0.0, 0.0}
	assert.Equal(t, expected, rotate(input, -1))
}
