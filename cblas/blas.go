// Do not manually edit this file. It was created by the genBlas.pl script from cblas.h.

// Copyright ©2012 The bíogo.blas Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cblas

/*
#cgo CFLAGS: -g -O2
#cgo LDFLAGS: -lcblas -latlas
#include "cblas.h"
*/
import "C"

import "unsafe"
import "github.com/dane-unltd/linalg/blas"

type BlasOps struct{}

func (ops BlasOps) Rotg(a float64, b float64) (c float64, s float64, r float64, z float64) {
	C.cblas_drotg((*C.double)(&a), (*C.double)(&b), (*C.double)(&c), (*C.double)(&s))
	return c, s, a, b
}
func (ops BlasOps) Rotmg(d1 float64, d2 float64, b1 float64, b2 float64) (p *blas.RotmParams, rd1 float64, rd2 float64, rb1 float64) {
	p = &blas.RotmParams{}
	C.cblas_drotmg((*C.double)(&d1), (*C.double)(&d2), (*C.double)(&b1), C.double(b2), (*C.double)(unsafe.Pointer(p)))
	return p, d1, d2, b1
}
func (ops BlasOps) Rotm(n int, x []float64, incX int, y []float64, incY int, p *blas.RotmParams) {
	C.cblas_drotm(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), (*C.double)(unsafe.Pointer(p)))
}

func (ops BlasOps) Dot(n int, x []float64, incX int, y []float64, incY int) float64 {
	return float64(C.cblas_ddot(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY)))
}
func (ops BlasOps) Nrm2(n int, x []float64, incX int) float64 {
	return float64(C.cblas_dnrm2(C.int(n), (*C.double)(&x[0]), C.int(incX)))
}
func (ops BlasOps) Asum(n int, x []float64, incX int) float64 {
	return float64(C.cblas_dasum(C.int(n), (*C.double)(&x[0]), C.int(incX)))
}
func (ops BlasOps) Iamax(n int, x []float64, incX int) int {
	return int(C.cblas_idamax(C.int(n), (*C.double)(&x[0]), C.int(incX)))
}
func (ops BlasOps) Swap(n int, x []float64, incX int, y []float64, incY int) {
	C.cblas_dswap(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY))
}
func (ops BlasOps) Copy(n int, x []float64, incX int, y []float64, incY int) {
	C.cblas_dcopy(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY))
}
func (ops BlasOps) Axpy(n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	C.cblas_daxpy(C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY))
}
func (ops BlasOps) Rot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	C.cblas_drot(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), C.double(c), C.double(s))
}
func (ops BlasOps) Scal(n int, alpha float64, x []float64, incX int) {
	C.cblas_dscal(C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX))
}
func (ops BlasOps) Gemv(o blas.Order, tA blas.Transpose, m int, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dgemv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (ops BlasOps) Gbmv(o blas.Order, tA blas.Transpose, m int, n int, kL int, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dgbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), C.int(kL), C.int(kU), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (ops BlasOps) Trmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtrmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX))
}
func (ops BlasOps) Tbmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX))
}
func (ops BlasOps) Tpmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
	C.cblas_dtpmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.double)(&ap[0]), (*C.double)(&x[0]), C.int(incX))
}
func (ops BlasOps) Trsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtrsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX))
}
func (ops BlasOps) Tbsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtbsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX))
}
func (ops BlasOps) Tpsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
	C.cblas_dtpsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.double)(&ap[0]), (*C.double)(&x[0]), C.int(incX))
}
func (ops BlasOps) Symv(o blas.Order, ul blas.Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dsymv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (ops BlasOps) Sbmv(o blas.Order, ul blas.Uplo, n int, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dsbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.int(k), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (ops BlasOps) Spmv(o blas.Order, ul blas.Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dspmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&ap[0]), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (ops BlasOps) Ger(o blas.Order, m int, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	C.cblas_dger(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), (*C.double)(&a[0]), C.int(lda))
}
func (ops BlasOps) Syr(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int) {
	C.cblas_dsyr(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&a[0]), C.int(lda))
}
func (ops BlasOps) Spr(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, ap []float64) {
	C.cblas_dspr(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&ap[0]))
}
func (ops BlasOps) Syr2(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	C.cblas_dsyr2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), (*C.double)(&a[0]), C.int(lda))
}
func (ops BlasOps) Spr2(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64) {
	C.cblas_dspr2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), (*C.double)(&a[0]))
}
func (ops BlasOps) Gemm(o blas.Order, tA blas.Transpose, tB blas.Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	C.cblas_dgemm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_TRANSPOSE(tB), C.int(m), C.int(n), C.int(k), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb), C.double(beta), (*C.double)(&c[0]), C.int(ldc))
}
func (ops BlasOps) Symm(o blas.Order, s blas.Side, ul blas.Uplo, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	C.cblas_dsymm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.int(m), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb), C.double(beta), (*C.double)(&c[0]), C.int(ldc))
}
func (ops BlasOps) Syrk(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int) {
	C.cblas_dsyrk(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), C.double(alpha), (*C.double)(&a[0]), C.int(lda), C.double(beta), (*C.double)(&c[0]), C.int(ldc))
}
func (ops BlasOps) Syr2k(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	C.cblas_dsyr2k(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb), C.double(beta), (*C.double)(&c[0]), C.int(ldc))
}
func (ops BlasOps) Trmm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	C.cblas_dtrmm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb))
}
func (ops BlasOps) Trsm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	C.cblas_dtrsm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb))
}
