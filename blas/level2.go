package blas

func Sgemv(o Order, tA Transpose, m int, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	Ops.Sgemv(o, tA, m, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Sgbmv(o Order, tA Transpose, m int, n int, kL int, kU int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	Ops.Sgbmv(o, tA, m, n, kL, kU, alpha, a, lda, x, incX, beta, y, incY)
}
func Strmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float32, lda int, x []float32, incX int) {
	Ops.Strmv(o, ul, tA, d, n, a, lda, x, incX)
}
func Stbmv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float32, lda int, x []float32, incX int) {
	Ops.Stbmv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Stpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float32, x []float32, incX int) {
	Ops.Stpmv(o, ul, tA, d, n, ap, x, incX)
}
func Strsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float32, lda int, x []float32, incX int) {
	Ops.Strsv(o, ul, tA, d, n, a, lda, x, incX)
}
func Stbsv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float32, lda int, x []float32, incX int) {
	Ops.Stbsv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Stpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float32, x []float32, incX int) {
	Ops.Stpsv(o, ul, tA, d, n, ap, x, incX)
}
func Dgemv(o Order, tA Transpose, m int, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	Ops.Dgemv(o, tA, m, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Dgbmv(o Order, tA Transpose, m int, n int, kL int, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	Ops.Dgbmv(o, tA, m, n, kL, kU, alpha, a, lda, x, incX, beta, y, incY)
}
func Dtrmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int) {
	Ops.Dtrmv(o, ul, tA, d, n, a, lda, x, incX)
}
func Dtbmv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	Ops.Dtbmv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Dtpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int) {
	Ops.Dtpmv(o, ul, tA, d, n, ap, x, incX)
}
func Dtrsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int) {
	Ops.Dtrsv(o, ul, tA, d, n, a, lda, x, incX)
}
func Dtbsv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	Ops.Dtbsv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Dtpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int) {
	Ops.Dtpsv(o, ul, tA, d, n, ap, x, incX)
}
func Cgemv(o Order, tA Transpose, m int, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	Ops.Cgemv(o, tA, m, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Cgbmv(o Order, tA Transpose, m int, n int, kL int, kU int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	Ops.Cgbmv(o, tA, m, n, kL, kU, alpha, a, lda, x, incX, beta, y, incY)
}
func Ctrmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	Ops.Ctrmv(o, ul, tA, d, n, a, lda, x, incX)
}
func Ctbmv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []complex64, lda int, x []complex64, incX int) {
	Ops.Ctbmv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Ctpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex64, x []complex64, incX int) {
	Ops.Ctpmv(o, ul, tA, d, n, ap, x, incX)
}
func Ctrsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	Ops.Ctrsv(o, ul, tA, d, n, a, lda, x, incX)
}
func Ctbsv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []complex64, lda int, x []complex64, incX int) {
	Ops.Ctbsv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Ctpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex64, x []complex64, incX int) {
	Ops.Ctpsv(o, ul, tA, d, n, ap, x, incX)
}
func Zgemv(o Order, tA Transpose, m int, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	Ops.Zgemv(o, tA, m, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Zgbmv(o Order, tA Transpose, m int, n int, kL int, kU int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	Ops.Zgbmv(o, tA, m, n, kL, kU, alpha, a, lda, x, incX, beta, y, incY)
}
func Ztrmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	Ops.Ztrmv(o, ul, tA, d, n, a, lda, x, incX)
}
func Ztbmv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []complex128, lda int, x []complex128, incX int) {
	Ops.Ztbmv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Ztpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex128, x []complex128, incX int) {
	Ops.Ztpmv(o, ul, tA, d, n, ap, x, incX)
}
func Ztrsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	Ops.Ztrsv(o, ul, tA, d, n, a, lda, x, incX)
}
func Ztbsv(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []complex128, lda int, x []complex128, incX int) {
	Ops.Ztbsv(o, ul, tA, d, n, k, a, lda, x, incX)
}
func Ztpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex128, x []complex128, incX int) {
	Ops.Ztpsv(o, ul, tA, d, n, ap, x, incX)
}
func Ssymv(o Order, ul Uplo, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	Ops.Ssymv(o, ul, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Ssbmv(o Order, ul Uplo, n int, k int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	Ops.Ssbmv(o, ul, n, k, alpha, a, lda, x, incX, beta, y, incY)
}
func Sspmv(o Order, ul Uplo, n int, alpha float32, ap []float32, x []float32, incX int, beta float32, y []float32, incY int) {
	Ops.Sspmv(o, ul, n, alpha, ap, x, incX, beta, y, incY)
}
func Sger(o Order, m int, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	Ops.Sger(o, m, n, alpha, x, incX, y, incY, a, lda)
}
func Ssyr(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, a []float32, lda int) {
	Ops.Ssyr(o, ul, n, alpha, x, incX, a, lda)
}
func Sspr(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, ap []float32) {
	Ops.Sspr(o, ul, n, alpha, x, incX, ap)
}
func Ssyr2(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	Ops.Ssyr2(o, ul, n, alpha, x, incX, y, incY, a, lda)
}
func Sspr2(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32) {
	Ops.Sspr2(o, ul, n, alpha, x, incX, y, incY, a)
}
func Dsymv(o Order, ul Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	Ops.Dsymv(o, ul, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Dsbmv(o Order, ul Uplo, n int, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	Ops.Dsbmv(o, ul, n, k, alpha, a, lda, x, incX, beta, y, incY)
}
func Dspmv(o Order, ul Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int) {
	Ops.Dspmv(o, ul, n, alpha, ap, x, incX, beta, y, incY)
}
func Dger(o Order, m int, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	Ops.Dger(o, m, n, alpha, x, incX, y, incY, a, lda)
}
func Dsyr(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int) {
	Ops.Dsyr(o, ul, n, alpha, x, incX, a, lda)
}
func Dspr(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, ap []float64) {
	Ops.Dspr(o, ul, n, alpha, x, incX, ap)
}
func Dsyr2(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	Ops.Dsyr2(o, ul, n, alpha, x, incX, y, incY, a, lda)
}
func Dspr2(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64) {
	Ops.Dspr2(o, ul, n, alpha, x, incX, y, incY, a)
}
func Chemv(o Order, ul Uplo, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	Ops.Chemv(o, ul, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Chbmv(o Order, ul Uplo, n int, k int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	Ops.Chbmv(o, ul, n, k, alpha, a, lda, x, incX, beta, y, incY)
}
func Chpmv(o Order, ul Uplo, n int, alpha complex64, ap []complex64, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	Ops.Chpmv(o, ul, n, alpha, ap, x, incX, beta, y, incY)
}
func Cgeru(o Order, m int, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	Ops.Cgeru(o, m, n, alpha, x, incX, y, incY, a, lda)
}
func Cgerc(o Order, m int, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	Ops.Cgerc(o, m, n, alpha, x, incX, y, incY, a, lda)
}
func Cher(o Order, ul Uplo, n int, alpha float32, x []complex64, incX int, a []complex64, lda int) {
	Ops.Cher(o, ul, n, alpha, x, incX, a, lda)
}
func Chpr(o Order, ul Uplo, n int, alpha float32, x []complex64, incX int, a []complex64) {
	Ops.Chpr(o, ul, n, alpha, x, incX, a)
}
func Cher2(o Order, ul Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	Ops.Cher2(o, ul, n, alpha, x, incX, y, incY, a, lda)
}
func Chpr2(o Order, ul Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, ap []complex64) {
	Ops.Chpr2(o, ul, n, alpha, x, incX, y, incY, ap)
}
func Zhemv(o Order, ul Uplo, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	Ops.Zhemv(o, ul, n, alpha, a, lda, x, incX, beta, y, incY)
}
func Zhbmv(o Order, ul Uplo, n int, k int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	Ops.Zhbmv(o, ul, n, k, alpha, a, lda, x, incX, beta, y, incY)
}
func Zhpmv(o Order, ul Uplo, n int, alpha complex128, ap []complex128, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	Ops.Zhpmv(o, ul, n, alpha, ap, x, incX, beta, y, incY)
}
func Zgeru(o Order, m int, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	Ops.Zgeru(o, m, n, alpha, x, incX, y, incY, a, lda)
}
func Zgerc(o Order, m int, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	Ops.Zgerc(o, m, n, alpha, x, incX, y, incY, a, lda)
}
func Zher(o Order, ul Uplo, n int, alpha float64, x []complex128, incX int, a []complex128, lda int) {
	Ops.Zher(o, ul, n, alpha, x, incX, a, lda)
}
func Zhpr(o Order, ul Uplo, n int, alpha float64, x []complex128, incX int, a []complex128) {
	Ops.Zhpr(o, ul, n, alpha, x, incX, a)
}
func Zher2(o Order, ul Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	Ops.Zher2(o, ul, n, alpha, x, incX, y, incY, a, lda)
}
func Zhpr2(o Order, ul Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, ap []complex128) {
	Ops.Zhpr2(o, ul, n, alpha, x, incX, y, incY, ap)
}
