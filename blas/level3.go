package blas

func Gemm(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	ops.Gemm(o, tA, tB, m, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Symm(o Order, s Side, ul Uplo, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	ops.Symm(o, s, ul, m, n, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Syrk(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int) {
	ops.Syrk(o, ul, t, n, k, alpha, a, lda, beta, c, ldc)
}
func Syr2k(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	ops.Syr2k(o, ul, t, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Trmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	ops.Trmm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
func Trsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	ops.Trsm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
