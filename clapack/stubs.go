package clapack

import (
	"github.com/dane-unltd/linalg/lapack"
	"github.com/gonum/blas"
)

func (Lapack) Dgees(jobvs lapack.Job, s lapack.Sort, sel lapack.Select, n int, a []float64, lda int, sdim int, wr []float64, wi []float64, vs []float64, ldvs int) lapack.Info {
	return 0
}

func (Lapack) Dgelqf(m int, n int, a []float64, lda int, tau []float64) lapack.Info {
	return 0
}
func (Lapack) Dgels(ta blas.Transpose, m int, n int, nrhs int, a []float64, lda int, b []float64, ldb int) lapack.Info {
	return 0
}
func (Lapack) Dgeqp3(m int, n int, a []float64, lda int, jpvt []int, tau []float64) lapack.Info {
	return 0
}
func (Lapack) Dgesdd(jobz lapack.Job, m int, n int, a []float64, lda int, s []float64, u []float64, ldu int, vt []float64, ldvt int) lapack.Info {
	return 0
}
func (Lapack) Dgges(jobvsl lapack.Job, jobvsr lapack.Job, s lapack.Sort, sel lapack.SelectG, n int, a []float64, lda int, b []float64, ldb int, sdim int, alphar []float64, alphai []float64, beta []float64, vsl []float64, ldvsl int, vsr []float64, ldvsr int) lapack.Info {
	return 0
}
func (Lapack) Dgtsv(n int, nrhs int, dl []float64, d []float64, du []float64, b []float64, ldb int) lapack.Info {
	return 0
}
func (Lapack) Dorglq(m int, n int, k int, a []float64, lda int, tau []float64) lapack.Info {
	return 0
}
func (Lapack) Dorgqr(m int, n int, k int, a []float64, lda int, tau []float64) lapack.Info {
	return 0
}
func (Lapack) Dormlq(s blas.Side, ta blas.Transpose, m int, n int, k int, a []float64, lda int, tau []float64, c []float64, ldc int) lapack.Info {
	return 0
}
func (Lapack) Dpbsv(ul blas.Uplo, n int, kd int, nrhs int, ab []float64, ldab int, b []float64, ldb int) lapack.Info {
	return 0
}
func (Lapack) Dpbtrf(ul blas.Uplo, n int, kd int, ab []float64, ldab int) lapack.Info {
	return 0
}
func (Lapack) Dpbtrs(ul blas.Uplo, n int, kd int, nrhs int, ab []float64, ldab int, b []float64, ldb int) lapack.Info {
	return 0
}
func (Lapack) Dptsv(n int, nrhs int, d []float64, e []float64, b []float64, ldb int) lapack.Info {
	return 0
}
func (Lapack) Dpttrf(n int, d []float64, e []float64) lapack.Info {
	return 0
}
func (Lapack) Dpttrs(n int, nrhs int, d []float64, e []float64, b []float64, ldb int) lapack.Info {
	return 0
}
func (Lapack) Dsyev(jobz lapack.Job, ul blas.Uplo, n int, a []float64, lda int, w []float64) lapack.Info {
	return 0
}
func (Lapack) Dsygv(t lapack.Type, jobz lapack.Job, ul blas.Uplo, n int, a []float64, lda int, b []float64, ldb int, w []float64) lapack.Info {
	return 0
}
func (Lapack) Dsysv(ul blas.Uplo, n int, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int) lapack.Info {
	return 0
}
func (Lapack) Dsytri(ul blas.Uplo, n int, a []float64, lda int, ipiv []int32) lapack.Info {
	return 0
}
func (Lapack) Dtbtrs(ul blas.Uplo, ta blas.Transpose, d blas.Diag, n int, kd int, nrhs int, ab []float64, ldab int, b []float64, ldb int) lapack.Info {
	return 0
}
func (Lapack) Dtrtri(ul blas.Uplo, d blas.Diag, n int, a []float64, lda int) lapack.Info {
	return 0
}
