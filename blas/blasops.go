package blas

type BlasOps struct {
	//level1
	Srotg  func(a float32, b float32) (c float32, s float32, r float32, z float32)
	Srotmg func(d1 float32, d2 float32, b1 float32, b2 float32) (p *SrotmParams, rd1 float32, rd2 float32, rb1 float32)
	Srotm  func(n int, x []float32, incX int, y []float32, incY int, p *SrotmParams)
	Drotg  func(a float64, b float64) (c float64, s float64, r float64, z float64)
	Drotmg func(d1 float64, d2 float64, b1 float64, b2 float64) (p *DrotmParams, rd1 float64, rd2 float64, rb1 float64)
	Drotm  func(n int, x []float64, incX int, y []float64, incY int, p *DrotmParams)

	Sdsdot func(n int, alpha float32, x []float32, incX int, y []float32, incY int) float32
	Dsdot  func(n int, x []float32, incX int, y []float32, incY int) float64
	Sdot   func(n int, x []float32, incX int, y []float32, incY int) float32
	Ddot   func(n int, x []float64, incX int, y []float64, incY int) float64
	Cdotu  func(n int, x []complex64, incX int, y []complex64, incY int) complex64
	Cdotc  func(n int, x []complex64, incX int, y []complex64, incY int) complex64
	Zdotu  func(n int, x []complex128, incX int, y []complex128, incY int) complex128
	Zdotc  func(n int, x []complex128, incX int, y []complex128, incY int) complex128

	Snrm2  func(n int, x []float32, incX int) float32
	Sasum  func(n int, x []float32, incX int) float32
	Dnrm2  func(n int, x []float64, incX int) float64
	Dasum  func(n int, x []float64, incX int) float64
	Scnrm2 func(n int, x []complex64, incX int) float32
	Scasum func(n int, x []complex64, incX int) float32
	Dznrm2 func(n int, x []complex128, incX int) float64
	Dzasum func(n int, x []complex128, incX int) float64
	Isamax func(n int, x []float32, incX int) int
	Idamax func(n int, x []float64, incX int) int
	Icamax func(n int, x []complex64, incX int) int
	Izamax func(n int, x []complex128, incX int) int
	Sswap  func(n int, x []float32, incX int, y []float32, incY int)
	Scopy  func(n int, x []float32, incX int, y []float32, incY int)
	Saxpy  func(n int, alpha float32, x []float32, incX int, y []float32, incY int)
	Dswap  func(n int, x []float64, incX int, y []float64, incY int)
	Dcopy  func(n int, x []float64, incX int, y []float64, incY int)
	Daxpy  func(n int, alpha float64, x []float64, incX int, y []float64, incY int)
	Cswap  func(n int, x []complex64, incX int, y []complex64, incY int)
	Ccopy  func(n int, x []complex64, incX int, y []complex64, incY int)
	Caxpy  func(n int, alpha complex64, x []complex64, incX int, y []complex64, incY int)
	Zswap  func(n int, x []complex128, incX int, y []complex128, incY int)
	Zcopy  func(n int, x []complex128, incX int, y []complex128, incY int)
	Zaxpy  func(n int, alpha complex128, x []complex128, incX int, y []complex128, incY int)
	Srot   func(n int, x []float32, incX int, y []float32, incY int, c float32, s float32)
	Drot   func(n int, x []float64, incX int, y []float64, incY int, c float64, s float64)
	Sscal  func(n int, alpha float32, x []float32, incX int)
	Dscal  func(n int, alpha float64, x []float64, incX int)
	Cscal  func(n int, alpha complex64, x []complex64, incX int)
	Zscal  func(n int, alpha complex128, x []complex128, incX int)
	Csscal func(n int, alpha float32, x []complex64, incX int)
	Zdscal func(n int, alpha float64, x []complex128, incX int)

	// level 2

	Sgemv func(o Order, tA Transpose, m int, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int)
	Sgbmv func(o Order, tA Transpose, m int, n int, kL int, kU int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int)
	Strmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float32, lda int, x []float32, incX int)
	Stbmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float32, lda int, x []float32, incX int)
	Stpmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float32, x []float32, incX int)
	Strsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float32, lda int, x []float32, incX int)
	Stbsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float32, lda int, x []float32, incX int)
	Stpsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float32, x []float32, incX int)
	Dgemv func(o Order, tA Transpose, m int, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Dgbmv func(o Order, tA Transpose, m int, n int, kL int, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Dtrmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int)
	Dtbmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float64, lda int, x []float64, incX int)
	Dtpmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int)
	Dtrsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int)
	Dtbsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []float64, lda int, x []float64, incX int)
	Dtpsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int)
	Cgemv func(o Order, tA Transpose, m int, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Cgbmv func(o Order, tA Transpose, m int, n int, kL int, kU int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Ctrmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex64, lda int, x []complex64, incX int)
	Ctbmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []complex64, lda int, x []complex64, incX int)
	Ctpmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex64, x []complex64, incX int)
	Ctrsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex64, lda int, x []complex64, incX int)
	Ctbsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []complex64, lda int, x []complex64, incX int)
	Ctpsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex64, x []complex64, incX int)
	Zgemv func(o Order, tA Transpose, m int, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Zgbmv func(o Order, tA Transpose, m int, n int, kL int, kU int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Ztrmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex128, lda int, x []complex128, incX int)
	Ztbmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []complex128, lda int, x []complex128, incX int)
	Ztpmv func(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex128, x []complex128, incX int)
	Ztrsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex128, lda int, x []complex128, incX int)
	Ztbsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, k int, a []complex128, lda int, x []complex128, incX int)
	Ztpsv func(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex128, x []complex128, incX int)
	Ssymv func(o Order, ul Uplo, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int)
	Ssbmv func(o Order, ul Uplo, n int, k int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int)
	Sspmv func(o Order, ul Uplo, n int, alpha float32, ap []float32, x []float32, incX int, beta float32, y []float32, incY int)
	Sger  func(o Order, m int, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int)
	Ssyr  func(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, a []float32, lda int)
	Sspr  func(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, ap []float32)
	Ssyr2 func(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int)
	Sspr2 func(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32)
	Dsymv func(o Order, ul Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Dsbmv func(o Order, ul Uplo, n int, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Dspmv func(o Order, ul Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int)
	Dger  func(o Order, m int, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int)
	Dsyr  func(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int)
	Dspr  func(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, ap []float64)
	Dsyr2 func(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int)
	Dspr2 func(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64)
	Chemv func(o Order, ul Uplo, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Chbmv func(o Order, ul Uplo, n int, k int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Chpmv func(o Order, ul Uplo, n int, alpha complex64, ap []complex64, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Cgeru func(o Order, m int, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int)
	Cgerc func(o Order, m int, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int)
	Cher  func(o Order, ul Uplo, n int, alpha float32, x []complex64, incX int, a []complex64, lda int)
	Chpr  func(o Order, ul Uplo, n int, alpha float32, x []complex64, incX int, a []complex64)
	Cher2 func(o Order, ul Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int)
	Chpr2 func(o Order, ul Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, ap []complex64)
	Zhemv func(o Order, ul Uplo, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Zhbmv func(o Order, ul Uplo, n int, k int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Zhpmv func(o Order, ul Uplo, n int, alpha complex128, ap []complex128, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Zgeru func(o Order, m int, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int)
	Zgerc func(o Order, m int, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int)
	Zher  func(o Order, ul Uplo, n int, alpha float64, x []complex128, incX int, a []complex128, lda int)
	Zhpr  func(o Order, ul Uplo, n int, alpha float64, x []complex128, incX int, a []complex128)
	Zher2 func(o Order, ul Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int)
	Zhpr2 func(o Order, ul Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, ap []complex128)

	//level 3

	Sgemm  func(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int)
	Ssymm  func(o Order, s Side, ul Uplo, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int)
	Ssyrk  func(o Order, ul Uplo, t Transpose, n int, k int, alpha float32, a []float32, lda int, beta float32, c []float32, ldc int)
	Ssyr2k func(o Order, ul Uplo, t Transpose, n int, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int)
	Strmm  func(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int)
	Strsm  func(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int)
	Dgemm  func(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
	Dsymm  func(o Order, s Side, ul Uplo, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
	Dsyrk  func(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int)
	Dsyr2k func(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
	Dtrmm  func(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int)
	Dtrsm  func(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int)
	Cgemm  func(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int)
	Csymm  func(o Order, s Side, ul Uplo, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int)
	Csyrk  func(o Order, ul Uplo, t Transpose, n int, k int, alpha complex64, a []complex64, lda int, beta complex64, c []complex64, ldc int)
	Csyr2k func(o Order, ul Uplo, t Transpose, n int, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int)
	Ctrmm  func(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int)
	Ctrsm  func(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int)
	Zgemm  func(o Order, tA Transpose, tB Transpose, m int, n int, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
	Zsymm  func(o Order, s Side, ul Uplo, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
	Zsyrk  func(o Order, ul Uplo, t Transpose, n int, k int, alpha complex128, a []complex128, lda int, beta complex128, c []complex128, ldc int)
	Zsyr2k func(o Order, ul Uplo, t Transpose, n int, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
	Ztrmm  func(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int)
	Ztrsm  func(o Order, s Side, ul Uplo, tA Transpose, d Diag, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int)
	Chemm  func(o Order, s Side, ul Uplo, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int)
	Cherk  func(o Order, ul Uplo, t Transpose, n int, k int, alpha float32, a []complex64, lda int, beta float32, c []complex64, ldc int)
	Cher2k func(o Order, ul Uplo, t Transpose, n int, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta float32, c []complex64, ldc int)
	Zhemm  func(o Order, s Side, ul Uplo, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
	Zherk  func(o Order, ul Uplo, t Transpose, n int, k int, alpha float64, a []complex128, lda int, beta float64, c []complex128, ldc int)
	Zher2k func(o Order, ul Uplo, t Transpose, n int, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta float64, c []complex128, ldc int)
}

var Ops BlasOps

func Register(ops BlasOps) {
	Ops = ops
}
