package matrix

var Ops struct {
	//level1
	Dnrm2  func(N int, X []float64, incX int) float64
	Dasum  func(N int, X []float64, incX int) float64
	Ddot   func(N int, X []float64, incX int, Y []float64, incY int) float64
	Idamax func(N int, X []float64, incX int) int
	Dswap  func(N int, X []float64, incX int, Y []float64, incY int)
	Dcopy  func(N int, X []float64, incX int, Y []float64, incY int)
	Daxpy  func(N int, alpha float64, X []float64, incX int, Y []float64, incY int)
	Dscal  func(N int, alpha float64, X []float64, incX int)
	//level 2
	Dgemv func(transA bool, M int, N int, alpha float64,
		A []float64, lda int, X []float64, incX int, beta float64,
		Y []float64, incY int)
	Dgbmv func(transA string, M int, N int, KL int, KU int,
		alpha float64, A []float64, lda int,
		X []float64, incX int, beta float64,
		Y []float64, incY int)
	Dtrmv func(uplo, transA, diag string,
		N int, A []float64, lda int, X []float64, incX int)
	Dtbmv func(uplo, transA, diag string,
		N int, K int, A []float64, lda int, X []float64, incX int)
	Dtpmv func(uplo, transA, diag string,
		N int, Ap []float64, X []float64, incX int)
	Dtrsv func(uplo string, transA bool, diag string,
		N int, A []float64, lda int, X []float64, incX int)
	Dtbsv func(uplo, transA, diag string,
		N int, K int, A []float64, lda int, X []float64, incX int)
	Dtpsv func(uplo, transA, diag string,
		N int, Ap []float64, X []float64, incX int)
	Dsymv func(uplo string, N int, alpha float64,
		A []float64, lda int, X []float64, incX int, beta float64,
		Y []float64, incY int)
	Dsbmv func(uplo string, N int, K int, alpha float64,
		A []float64, lda int, X []float64, incX int, beta float64,
		Y []float64, incY int)
	Dspmv func(uplo string, N int, alpha float64,
		Ap []float64, X []float64, incX int, beta float64,
		Y []float64, incY int)
	Dger func(M int, N int, alpha float64,
		X []float64, incX int, Y []float64, incY int,
		A []float64, lda int)
	Dsyr func(uplo string, N int, alpha float64,
		X []float64, incX int, A []float64, lda int)
	Dspr func(uplo string, N int, alpha float64,
		X []float64, incX int, Ap []float64)
	Dsyr2 func(uplo string, N int, alpha float64,
		X []float64, incX int, Y []float64, incY int, A []float64, lda int)
	Dspr2 func(uplo string, N int, alpha float64,
		X []float64, incX int, Y []float64, incY int, Ap []float64)
	//level3
	Dgemm func(tA, tB bool, m int, n int, k int,
		alpha float64, a []float64, lda int, b []float64, ldb int,
		beta float64, c []float64, ldc int)
	Dsymm func(side, uplo string, M int, N int,
		alpha float64, A []float64, lda int, B []float64, ldb int, beta float64,
		C []float64, ldc int)
	Dsyrk func(uplo, trans string, N int, K int,
		alpha float64, A []float64, lda int, beta float64,
		C []float64, ldc int)
	Dsyr2k func(uplo, trans string, N int, K int,
		alpha float64, A []float64, lda int, B []float64, ldb int, beta float64,
		C []float64, ldc int)
	Dtrmm func(side, uplo, transA, diag string,
		M int, N int, alpha float64, A []float64, lda int, B []float64, ldb int)
	Dtrsm func(side, uplo, transA, diag string,
		M int, N int, alpha float64, A []float64, lda int, B []float64, ldb int)

	//lapack
	Dgesvd func(jobu, jobvt string, M, N int, A []float64, lda int, S []float64, U []float64,
		ldu int, Vt []float64, ldvt int) int
	Dgeqrf func(M, N int, A []float64, lda int, tau []float64) int
	Dpotrf func(uplo string, N int, A []float64, lda int) int
	Dsyevd func(jobz, uplo string, N int, A []float64, lda int, W []float64) int
}
