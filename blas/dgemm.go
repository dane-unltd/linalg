package blas

import "github.com/dane-unltd/linalg/blasops"

func Dgemm(o blasops.Order, tA, tB blasops.Transpose, m int, n int, k int,
	alpha float64, a []float64, lda int, b []float64, ldb int,
	beta float64, c []float64, ldc int) {

	var inner, outer, veclen int
	if o == blasops.ColMajor {
		outer = n
		veclen = m
	} else {
		veclen = n
		outer = m
		a, b = b, a
		ldb, lda = lda, ldb
		tA, tB = tB, tA
	}
	inner = k

	for j := 0; j < outer; j++ {
		cj := c[j*ldc:]
		Dscal(veclen, beta, cj, 1)

		for l := 0; l < inner; l++ {
			al := a[l*lda:]
			if tA == blasops.Trans {
				al = a[l:]
			}
			blj := b[j*ldb+l]
			if tB == blasops.Trans {
				blj = b[l*ldb+j]
			}
			if tA == blasops.NoTrans {
				Daxpy(veclen, blj*alpha, al, 1, cj, 1)
			} else {
				Daxpy(veclen, blj*alpha, al, lda, cj, 1)
			}
		}
	}
}
