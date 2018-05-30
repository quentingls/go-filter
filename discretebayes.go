package gofilter

// Predict performs the prediction step of the discrete bayes filter
func Predict(belief []float64, kernel []float64, offset int) []float64 {
	return convolute1d(rotate(belief, offset), kernel)
}

// Update performs an update on the belief of a discrete distribution
func Update(likelihood, prior []float64) []float64 {
	posterior := make([]float64, len(prior))
	sum := 0.0
	for i, _ := range prior {
		posterior[i] = prior[i] * likelihood[i]
		sum += posterior[i]
	}
	for i, _ := range posterior {
		posterior[i] /= sum
	}
	return posterior
}

// Uni-dimensional convolution with the kernel centered on the point.
func convolute1d(input []float64, kernel []float64) []float64 {
	output := make([]float64, len(input))
	offset := len(kernel) / 2

	for i, _ := range input {
		for j, ker := range kernel {
			idx := (i + offset - j + len(input)) % len(input)
			output[i] += input[idx] * ker
		}
	}
	return output
}
