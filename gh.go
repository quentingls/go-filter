package gofilter

// TODO: Create GH filter struct with update loop
func GHFilter(measurements []float64, x0, dx0, g, h, dt float64) []float64 {
	estimations := make([]float64, len(measurements))
	x_hat := x0
	dx_hat := dx0
	var residual float64

	for i, measurement := range measurements {
		// prediction step
		prediction := x_hat + dx_hat*dt // constant speed
		dx_hat = dx_hat                 // no accelration

		// update
		residual = measurement - prediction
		dx_hat = dx_hat + h*residual/dt // update the speed
		x_hat = prediction + g*residual
		estimations[i] = x_hat
	}

	return estimations
}
