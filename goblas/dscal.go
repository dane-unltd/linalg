package goblas

func (Blas) Dscal(N int, alpha float64, X []float64, incX int) {
	var xi int
	for ; N >= 2; N -= 2 {
		X[xi] = alpha * X[xi]
		xi += incX
		X[xi] = alpha * X[xi]
		xi += incX
	}
	if N != 0 {
		X[xi] = alpha * X[xi]
	}
}
