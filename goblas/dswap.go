package goblas

func (Blas) Dswap(N int, X []float64, incX int, Y []float64, incY int) {
	var xi, yi int
	for ; N > 0; N-- {
		X[xi], Y[yi] = Y[yi], X[xi]
		xi += incX
		yi += incY
	}
}
