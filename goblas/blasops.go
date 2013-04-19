package goblas

import "github.com/dane-unltd/linalg/blas"

type BlasOps struct{}

//level1
func (op BlasOps) Rotg(a float64, b float64) (c float64, s float64, r float64, z float64) {
	return Drotg(a, b)
}
func (op BlasOps) Rotmg(d1 float64, d2 float64, b1 float64, b2 float64) (p *blas.RotmParams, rd1 float64, rd2 float64, rb1 float64) {
	return Drotmg(d1, d2, b1, b2)
}
func (op BlasOps) Rotm(n int, x []float64, incX int, y []float64, incY int, p *blas.RotmParams) {
}

func (op BlasOps) Dot(n int, x []float64, incX int, y []float64, incY int) float64 {
	return Ddot(n, x, incX, y, incY)
}
func (op BlasOps) Nrm2(n int, x []float64, incX int) float64 {
	return Dnrm2(n, x, incX)
}
func (op BlasOps) Asum(n int, x []float64, incX int) float64 {
	return Dasum(n, x, incX)
}
func (op BlasOps) Iamax(n int, x []float64, incX int) int {
	return Idamax(n, x, incX)
}
func (op BlasOps) Swap(n int, x []float64, incX int, y []float64, incY int) {
	Dswap(n, x, incX, y, incY)
}
func (op BlasOps) Copy(n int, x []float64, incX int, y []float64, incY int) {
	Dcopy(n, x, incX, y, incY)
}
func (op BlasOps) Axpy(n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	Daxpy(n, alpha, x, incX, y, incY)
}
func (op BlasOps) Rot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	Drot(n, x, incX, y, incY, c, s)
}
func (op BlasOps) Scal(n int, alpha float64, x []float64, incX int) {
	Dscal(n, alpha, x, incX)
}

// level 2

func (op BlasOps) Gemv(o blas.Order, tA blas.Transpose, m int, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
}
func (op BlasOps) Gbmv(o blas.Order, tA blas.Transpose, m int, n int, kL int, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
}
func (op BlasOps) Trmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
}
func (op BlasOps) Tbmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
}
func (op BlasOps) Tpmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
}
func (op BlasOps) Trsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
}
func (op BlasOps) Tbsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
}
func (op BlasOps) Tpsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
}
func (op BlasOps) Symv(o blas.Order, ul blas.Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
}
func (op BlasOps) Sbmv(o blas.Order, ul blas.Uplo, n int, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
}
func (op BlasOps) Spmv(o blas.Order, ul blas.Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int) {
}
func (op BlasOps) Ger(o blas.Order, m int, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
}
func (op BlasOps) Syr(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int) {
}
func (op BlasOps) Spr(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, ap []float64) {
}
func (op BlasOps) Syr2(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
}
func (op BlasOps) Spr2(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64) {
}

//level 3

func (op BlasOps) Gemm(o blas.Order, tA blas.Transpose, tB blas.Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	Dgemm(o, tA, tB, m, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func (op BlasOps) Symm(o blas.Order, s blas.Side, ul blas.Uplo, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
}
func (op BlasOps) Syrk(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int) {
}
func (op BlasOps) Syr2k(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
}
func (op BlasOps) Trmm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
}
func (op BlasOps) Trsm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
}
