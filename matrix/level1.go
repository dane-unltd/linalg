package matrix

import "math"

var Ddot = func(X []float64, Y []float64) float64 {
	res := 0.0
	for i := range X {
		res += X[i] * Y[i]
	}
	return res
}

var Dnrm2 = func(X []float64) float64 {
	res := 0.0
	for _, v := range X {
		res += v * v
	}
	return math.Sqrt(res)
}
