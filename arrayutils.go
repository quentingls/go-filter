package gofilter

import (
	"math"
)

// Right rotation
func rotate(slice []float64, rotation int) []float64 {
	rotated := make([]float64, len(slice))
	for i, elt := range slice {
		rotated[(i+len(slice)+rotation)%len(slice)] = elt
	}
	return rotated
}

func roundArray(inputs []float64, places int64) []float64 {
	outputs := make([]float64, len(inputs))
	exp := math.Pow(10, float64(places))
	var round, val, remaining float64
	for i, input := range inputs {
		val = exp * input
		_, remaining = math.Modf(val)
		if remaining >= 0.5 {
			round = math.Ceil(val)
		} else {
			round = math.Floor(val)
		}
		outputs[i] = round / exp
	}
	return outputs
}
