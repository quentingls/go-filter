package gofilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGHFilter(t *testing.T) {
	a := assert.New(t)

	weights := []float64{158.0, 164.2, 160.3, 159.9, 162.1, 164.6,
		169.6, 167.4, 166.4, 171.0, 171.2, 172.6}
	expectedEstimations := []float64{159.2, 161.8, 162.1, 160.78, 160.985, 163.311, 168.1, 169.696,
		168.204, 169.164, 170.892, 172.629}
	actualEstimations := GHFilter(weights, 160.0, 1.0, 0.6, 2.0/3.0, 1.0)
	a.Equal(expectedEstimations, roundArray(actualEstimations, 3))
}
