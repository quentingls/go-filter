package gofilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPredict(t *testing.T) {
	belief := []float64{0.0, 0.0, 0.0, 0.2, 0.7, 0.1, 0.0, 0.0, 0.0, 0.0}
	expectedPrior := []float64{0.0,0.0,0.0, 0.0, 0.02, 0.23, 0.59, 0.15, 0.01, 0.0}
	kernel := []float64{0.1, 0.8, 0.1}
	assert.Equal(t, expectedPrior, roundArray(Predict(belief, kernel, 2), 2))
}

func TestUpdate(t *testing.T) {
	prior := []float64{0.2, 0.2, 0.2, 0.2, 0.2}
	likelihood := []float64{3, 1, 3, 1, 3}

	expectedPosterior := []float64{0.27, 0.09, 0.27, 0.09, 0.27}
	actualPosterior := roundArray(Update(likelihood, prior), 2)
	assert.Equal(t, expectedPosterior, actualPosterior)
}

func TestConvolute(t *testing.T) {
	input := []float64{0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
	kernel := []float64{0.1, 0.8, 0.1}
	expected := []float64{0.0, 0.0, 0.1, 0.8, 0.1, 0.0, 0.0, 0.0, 0.0, 0.0}
	//assert.Equal(t, expected, convolute1d(input, kernel))
	kernel = []float64{0.1, 0.9}
	expected = []float64{0.0, 0.0, 0.1, 0.9, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
	//assert.Equal(t, expected, convolute1d(input, kernel))
	kernel = []float64{0.1, 0.7, 0.2}
	expected = []float64{0.0, 0.0, 0.1, 0.7, 0.2, 0.0, 0.0, 0.0, 0.0, 0.0}
	assert.Equal(t, expected, convolute1d(input, kernel))
}
