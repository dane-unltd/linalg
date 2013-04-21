package goblas

import "github.com/kortschak/blas"

// Performs: y = alpha * A * x + beta * y  or y = alpha * A^T * x + beta * y
func (Blas) Dgemv(o blas.Order, tA blas.Transpose, m, n int, alpha float64, a []float64,
	lda int, x []float64, incX int, beta float64, y []float64, incY int) {

	dscal(m, beta, y, incY)
	if (o == blas.ColMajor && tA == blas.NoTrans) ||
		(o == blas.RowMajor && tA == blas.Trans) {
		for c := 0; c < n; c++ {
			daxpy(m, alpha*x[incX*c], a[c*lda:], 1, y, incY)
		}
	} else {
		for r := 0; r < m; r++ {
			y[r*incY] += alpha * ddot(n, a[r*lda:], 1, x, incX)
		}
	}

}
