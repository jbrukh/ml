package ml

// SimpleRegression returns the parameters of the linear
// model that best fits the given data. x and y must be the
// same size and contain at least two observations. theta0
// is the y-intercept and theta1 is the slope of the
// resulting line.
func SimpleLinearRegression(x []float64, y []float64) (theta0, theta1 float64) {
	N := len(x)
    if N != len(y) || N < 2 {
        panic("must have at least two data points")
    }
	meanOfX := sum(x)/float64(N)
	meanOfY := sum(y)/float64(N)
	cv, vr := float64(0), float64(0)
	for i := 0; i < N; i++ {
		xNorm := (x[i]-meanOfX)
		cv += xNorm*(y[i]-meanOfY)
		vr += xNorm*xNorm
	}
	theta1 = cv/vr
	theta0 = (meanOfY - theta1*meanOfX)
	return
}

func sum(arr []float64) (result float64) {
	for _, x := range arr {
		result += x
	}
	return
}
