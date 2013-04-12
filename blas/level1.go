package blas

import "math"

var Ddot = func(n int, x []float64, incX int, y []float64, incY int) float64 {
	res := 0.0
	for i := 0; i < n; i++ {
		res += x[i*incX] * y[i*incY]
	}
	return res
}

var Dnrm2 = func(n int, x []float64, incX int) float64 {
	res := 0.0
	v := 0.0
	for i := 0; i < n; i++ {
		v = x[i*incX]
		res += v * v
	}
	return math.Sqrt(res)
}

var Daxpy = func(N int, alpha float64, X []float64, incX int, Y []float64, incY int) {
	var xi, yi int
	switch alpha {
	case 0:
	case 1:
		for ; N >= 2; N -= 2 {
			Y[yi] += X[xi]
			xi += incX
			yi += incY

			Y[yi] += X[xi]
			xi += incX
			yi += incY
		}
		if N != 0 {
			Y[yi] += alpha * X[xi]
		}
	case -1:
		for ; N >= 2; N -= 2 {
			Y[yi] -= X[xi]
			xi += incX
			yi += incY

			Y[yi] -= X[xi]
			xi += incX
			yi += incY
		}
		if N != 0 {
			Y[yi] -= X[xi]
		}
	default:
		for ; N >= 2; N -= 2 {
			Y[yi] += alpha * X[xi]
			xi += incX
			yi += incY

			Y[yi] += alpha * X[xi]
			xi += incX
			yi += incY
		}
		if N != 0 {
			Y[yi] += alpha * X[xi]
		}
	}
}

var Dasum = func(n int, x []float64, incX int) float64 {
	res := 0.0
	for i := 0; i < n; i++ {
		res += math.Abs(x[i*incX])
	}
	return res
}

var Dcopy = func(n int, x []float64, incX int, y []float64, incY int) {
	for i := 0; i < n; i++ {
		y[i*incY] = x[i*incX]
	}
}

var Dscal = func(n int, alpha float64, x []float64, incX int) {
	for i := 0; i < n; i++ {
		x[i*incX] *= alpha
	}
}
