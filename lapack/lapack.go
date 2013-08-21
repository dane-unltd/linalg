package lapack

import (
	"github.com/gonum/blas"
)

type Info int
type Job byte

type Float64 interface {
	Dgesvd(order blas.Order, jobu Job, jobvt Job, m int, n int, a []float64,
		lda int, s []float64, u []float64, ldu int, vt []float64,
		ldvt int, superb []float64) Info
	Dpotrf(order blas.Order, uplo blas.Uplo, n int, a []float64,
		lda int) Info
	Dsytrf(order blas.Order, uplo blas.Uplo, n int, a []float64,
		lda int, ipiv []int) Info
}
