package blas

type BlasOps interface {
	//level1
	Rotg(a float64, b float64) (c float64, s float64, r float64, z float64)
	Rotmg(d1 float64, d2 float64, b1 float64, b2 float64) (p *RotmParams, rd1 float64, rd2 float64, rb1 float64)
	Rotm(n int, x []float64, incX int, y []float64, incY int, p *RotmParams)

	Dot(n int, x []float64, incX int, y []float64, incY int) float64
	Nrm2(n int, x []float64, incX int) float64
	Asum(n int, x []float64, incX int) float64
	Iamax(n int, x []float64, incX int) int
	Swap(n int, x []float64, incX int, y []float64, incY int)
	Copy(n int, x []float64, incX int, y []float64, incY int)
	Axpy(n int, alpha float64, x []float64, incX int, y []float64, incY int)
	Rot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64)
	Scal(n int, alpha float64, x []float64, incX int)

	// level 2

	Gemv(o Order, tA Transpose, m int, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Gbmv(o Order, tA Transpose, m int, n int, kL int, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Trmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int)
	Tbmv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float64, lda int, x []float64, incX int)
	Tpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int)
	Trsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int)
	Tbsv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float64, lda int, x []float64, incX int)
	Tpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int)
	Symv(o Order, ul Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Sbmv(o Order, ul Uplo, n int, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Spmv(o Order, ul Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int)
	Ger(o Order, m int, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int)
	Syr(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int)
	Spr(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, ap []float64)
	Syr2(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int)
	Spr2(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64)

	//level 3

	Gemm(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
	Symm(o Order, s Side, ul Uplo, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
	Syrk(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int)
	Syr2k(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
	Trmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int)
	Trsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int)
}

var ops BlasOps

func Register(o BlasOps) {
	ops = ops
}
