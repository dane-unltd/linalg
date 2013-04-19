// Copyright ©2012 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package blas provides bindings to the CBLAS linear algebra library.

Quick Reference Guide to the BLAS from http://www.netlib.org/lapack/lug/node145.html

Level 1 BLAS

	        dim scalar vector   vector   scalars              5-element prefixes
	                                                          struct

	_rotg (                                      a, b )                S, D (and C, Z if ATLAS included)
	_rotmg(                              d1, d2, a, b )                S, D
	_rot  ( n,         x, incX, y, incY,               c, s )          S, D
	_rotm ( n,         x, incX, y, incY,                      param )  S, D
	_swap ( n,         x, incX, y, incY )                              S, D, C, Z
	_scal ( n,  alpha, x, incX )                                       S, D, C, Z, Cs, Zd
	_copy ( n,         x, incX, y, incY )                              S, D, C, Z
	_axpy ( n,  alpha, x, incX, y, incY )                              S, D, C, Z
	_dot  ( n,         x, incX, y, incY )                              S, D, Ds
	_dotu ( n,         x, incX, y, incY )                              C, Z
	_dotc ( n,         x, incX, y, incY )                              C, Z
	__dot ( n,  alpha, x, incX, y, incY )                              Sds
	_nrm2 ( n,         x, incX )                                       S, D, Sc, Dz
	_asum ( n,         x, incX )                                       S, D, Sc, Dz
	I_amax( n,         x, incX )                                       s, d, c, z

Level 2 BLAS

	        options                   dim   b-width scalar matrix  vector   scalar vector   prefixes

	_gemv ( order,        trans,      m, n,         alpha, a, lda, x, incX, beta,  y, incY ) S, D, C, Z
	_gbmv ( order,        trans,      m, n, kL, kU, alpha, a, lda, x, incX, beta,  y, incY ) S, D, C, Z
	_hemv ( order, uplo,                 n,         alpha, a, lda, x, incX, beta,  y, incY ) C, Z
	_hbmv ( order, uplo,                 n, k,      alpha, a, lda, x, incX, beta,  y, incY ) C, Z
	_hpmv ( order, uplo,                 n,         alpha, ap,     x, incX, beta,  y, incY ) C, Z
	_symv ( order, uplo,                 n,         alpha, a, lda, x, incX, beta,  y, incY ) S, D
	_sbmv ( order, uplo,                 n, k,      alpha, a, lda, x, incX, beta,  y, incY ) S, D
	_spmv ( order, uplo,                 n,         alpha, ap,     x, incX, beta,  y, incY ) S, D
	_trmv ( order, uplo, trans, diag,    n,                a, lda, x, incX )                 S, D, C, Z
	_tbmv ( order, uplo, trans, diag,    n, k,             a, lda, x, incX )                 S, D, C, Z
	_tpmv ( order, uplo, trans, diag,    n,                ap,     x, incX )                 S, D, C, Z
	_trsv ( order, uplo, trans, diag,    n,                a, lda, x, incX )                 S, D, C, Z
	_tbsv ( order, uplo, trans, diag,    n, k,             a, lda, x, incX )                 S, D, C, Z
	_tpsv ( order, uplo, trans, diag,    n,                ap,     x, incX )                 S, D, C, Z

	        options                   dim   scalar vector   vector   matrix  prefixes

	_ger  ( order,                    m, n, alpha, x, incX, y, incY, a, lda ) S, D
	_geru ( order,                    m, n, alpha, x, incX, y, incY, a, lda ) C, Z
	_gerc ( order,                    m, n, alpha, x, incX, y, incY, a, lda ) C, Z
	_her  ( order, uplo,                 n, alpha, x, incX,          a, lda ) C, Z
	_hpr  ( order, uplo,                 n, alpha, x, incX,          ap )     C, Z
	_her2 ( order, uplo,                 n, alpha, x, incX, y, incY, a, lda ) C, Z
	_hpr2 ( order, uplo,                 n, alpha, x, incX, y, incY, ap )     C, Z
	_syr  ( order, uplo,                 n, alpha, x, incX,          a, lda ) S, D
	_spr  ( order, uplo,                 n, alpha, x, incX,          ap )     S, D
	_syr2 ( order, uplo,                 n, alpha, x, incX, y, incY, a, lda ) S, D
	_spr2 ( order, uplo,                 n, alpha, x, incX, y, incY, ap )     S, D

Level 3 BLAS

	        options                                 dim      scalar matrix  matrix  scalar matrix  prefixes

	_gemm ( order,             transA, transB,      m, n, k, alpha, a, lda, b, ldb, beta,  c, ldc ) S, D, C, Z
	_symm ( order, side, uplo,                      m, n,    alpha, a, lda, b, ldb, beta,  c, ldc ) S, D, C, Z
	_hemm ( order, side, uplo,                      m, n,    alpha, a, lda, b, ldb, beta,  c, ldc ) C, Z
	_syrk ( order,       uplo, trans,                  n, k, alpha, a, lda,         beta,  c, ldc ) S, D, C, Z
	_herk ( order,       uplo, trans,                  n, k, alpha, a, lda,         beta,  c, ldc ) C, Z
	_syr2k( order,       uplo, trans,                  n, k, alpha, a, lda, b, ldb, beta,  c, ldc ) S, D, C, Z
	_her2k( order,       uplo, trans,                  n, k, alpha, a, lda, b, ldb, beta,  c, ldc ) C, Z
	_trmm ( order, side, uplo, transA,        diag, m, n,    alpha, a, lda, b, ldb )                S, D, C, Z
	_trsm ( order, side, uplo, transA,        diag, m, n,    alpha, a, lda, b, ldb )                S, D, C, Z

Meaning of prefixes

	S - float32	C - complex64 	
	D - float64	Z - complex128 

Matrix types

	GE - GEneral 		GB - General Band 	
	SY - SYmmetric 		SB - Symmetric Band 	SP - Symmetric Packed
	HE - HErmitian 		HB - Hermitian Band 	HP - Hermitian Packed
	TR - TRiangular 	TB - Triangular Band 	TP - Triangular Packed

Options

	trans 	= CblasNoTrans, CblasTrans, CblasConjTrans
	uplo 	= CblasUpper, CblasLower
	diag 	= CblasNonunit, CblasUnit
	side 	= CblasLeft, CblasRight (A or op(A) on the left, or A or op(A) on the right)

For real matrices, CblasTrans and CblasConjTrans have the same meaning.
For Hermitian matrices, trans = CblasTrans is not allowed.
For complex symmetric matrices, trans = CblasConjTrans is not allowed.
*/
package cblas
