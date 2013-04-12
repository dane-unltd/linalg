package blas

var Dgemm = func(o, tA, tB int, m int, n int, k int,
	alpha float64, a []float64, lda int, b []float64, ldb int,
	beta float64, c []float64, ldc int) {

	for j := 0; j < n; j++ {
		cj := c[j*ldc : j*ldc+m]
		for l := 0; l < k; l++ {
			al := a[l*lda : l*lda+m]
			blj := b[j*ldb+l]
			Daxpy(m, blj, al, 1, cj, 1)
		}
	}
}
