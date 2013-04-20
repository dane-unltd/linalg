package goblas

import "github.com/kortschak/blas"

// Performs: y = alpha * A * x + beta * y  or y = alpha * A^T * x + beta * y
func (Blas) Dgemv(o blas.Order, tA blas.Transpose, m, n int, alpha float64, a []float64,
	lda int, x []float64, incX int, beta float64, y []float64, incY int) {

}
