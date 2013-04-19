package blas

func Gemv(o Order, tA Transpose, m int, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	ops.Gemv(o, tA, m, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Gbmv(o Order, tA Transpose, m int, n int, kL int, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	ops.Gbmv(o, tA, m, n, kL, kU, alpha, a, lda, x, incX, beta, y, incY)
}
func Trmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int) {
	ops.Trmv(o, ul, tA, d, n, a, lda, x, incX)
}
func Tbmv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	ops.Tbmv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Tpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int) {
	ops.Tpmv(o, ul, tA, d, n, ap, x, incX)
}
func Trsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int) {
	ops.Trsv(o, ul, tA, d, n, a, lda, x, incX)
}
func Tbsv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	ops.Tbsv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Tpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int) {
	ops.Tpsv(o, ul, tA, d, n, ap, x, incX)
}
func Symv(o Order, ul Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	ops.Symv(o, ul, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Sbmv(o Order, ul Uplo, n int, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	ops.Sbmv(o, ul, n, k, alpha, a, lda, x, incX, beta, y, incY)
}
func Spmv(o Order, ul Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int) {
	ops.Spmv(o, ul, n, alpha, ap, x, incX, beta, y, incY)
}
func Ger(o Order, m int, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	ops.Ger(o, m, n, alpha, x, incX, y, incY, a, lda)
}
func Syr(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int) {
	ops.Syr(o, ul, n, alpha, x, incX, a, lda)
}
func Spr(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, ap []float64) {
	ops.Spr(o, ul, n, alpha, x, incX, ap)
}
func Syr2(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	ops.Syr2(o, ul, n, alpha, x, incX, y, incY, a, lda)
}
func Spr2(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64) {
	ops.Spr2(o, ul, n, alpha, x, incX, y, incY, a)
}
