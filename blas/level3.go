package blas

func Sgemm(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	Ops.Sgemm(o, tA, tB, m, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Ssymm(o Order, s Side, ul Uplo, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	Ops.Ssymm(o, s, ul, m, n, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Ssyrk(o Order, ul Uplo, t Transpose, n int, k int, alpha float32, a []float32, lda int, beta float32, c []float32, ldc int) {
	Ops.Ssyrk(o, ul, t, n, k, alpha, a, lda, beta, c, ldc)
}
func Ssyr2k(o Order, ul Uplo, t Transpose, n int, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	Ops.Ssyr2k(o, ul, t, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Strmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int) {
	Ops.Strmm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
func Strsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int) {
	Ops.Strsm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
func Dgemm(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	Ops.Dgemm(o, tA, tB, m, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Dsymm(o Order, s Side, ul Uplo, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	Ops.Dsymm(o, s, ul, m, n, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Dsyrk(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int) {
	Ops.Dsyrk(o, ul, t, n, k, alpha, a, lda, beta, c, ldc)
}
func Dsyr2k(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	Ops.Dsyr2k(o, ul, t, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Dtrmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	Ops.Dtrmm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
func Dtrsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	Ops.Dtrsm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
func Cgemm(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	Ops.Cgemm(o, tA, tB, m, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Csymm(o Order, s Side, ul Uplo, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	Ops.Csymm(o, s, ul, m, n, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Csyrk(o Order, ul Uplo, t Transpose, n int, k int, alpha complex64, a []complex64, lda int, beta complex64, c []complex64, ldc int) {
	Ops.Csyrk(o, ul, t, n, k, alpha, a, lda, beta, c, ldc)
}
func Csyr2k(o Order, ul Uplo, t Transpose, n int, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	Ops.Csyr2k(o, ul, t, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Ctrmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int) {
	Ops.Ctrmm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
func Ctrsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int) {
	Ops.Ctrsm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
func Zgemm(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	Ops.Zgemm(o, tA, tB, m, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Zsymm(o Order, s Side, ul Uplo, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	Ops.Zsymm(o, s, ul, m, n, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Zsyrk(o Order, ul Uplo, t Transpose, n int, k int, alpha complex128, a []complex128, lda int, beta complex128, c []complex128, ldc int) {
	Ops.Zsyrk(o, ul, t, n, k, alpha, a, lda, beta, c, ldc)
}
func Zsyr2k(o Order, ul Uplo, t Transpose, n int, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	Ops.Zsyr2k(o, ul, t, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Ztrmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int) {
	Ops.Ztrmm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
func Ztrsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int) {
	Ops.Ztrsm(o, s, ul, tA, d, m, n, alpha, a, lda, b, ldb)
}
func Chemm(o Order, s Side, ul Uplo, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	Ops.Chemm(o, s, ul, m, n, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Cherk(o Order, ul Uplo, t Transpose, n int, k int, alpha float32, a []complex64, lda int, beta float32, c []complex64, ldc int) {
	Ops.Cherk(o, ul, t, n, k, alpha, a, lda, beta, c, ldc)
}
func Cher2k(o Order, ul Uplo, t Transpose, n int, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta float32, c []complex64, ldc int) {
	Ops.Cher2k(o, ul, t, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Zhemm(o Order, s Side, ul Uplo, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	Ops.Zhemm(o, s, ul, m, n, alpha, a, lda, b, ldb, beta, c, ldc)
}
func Zherk(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []complex128, lda int, beta float64, c []complex128, ldc int) {
	Ops.Zherk(o, ul, t, n, k, alpha, a, lda, beta, c, ldc)
}
func Zher2k(o Order, ul Uplo, t Transpose, n int, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta float64, c []complex128, ldc int) {
	Ops.Zher2k(o, ul, t, n, k, alpha, a, lda, b, ldb, beta, c, ldc)
}
