// Copyright (c) Harri Rautila, 2012

// This file is part of go.opt/linalg/lapack package.
// It is free software, distributed under the terms of GNU Lesser General Public
// License Version 3, or any later version. See the COPYING tile included in this archive.

package lapacke

// #cgo linux LDFLAGS: -lopenblas
// #cgo darwin LDFLAGS: -L/usr/local/Cellar/openblas/0.2.6/lib -lopenblas
// #include <stdlib.h>
// #include "lapacke.h"
import "C"
import (
	"github.com/dane-unltd/linalg/lapack"
	"github.com/gonum/blas"
	"unsafe"
)

//import "fmt"

type Lapack struct{}

var (
	_ lapack.Float64 = Lapack{}
)

func procUplo(uplo blas.Uplo) C.char {
	if uplo == blas.Upper {
		return 'U'
	}
	if uplo == blas.Lower {
		return 'L'
	}
	return 0
}

func (Lapack) Dgesvd(order blas.Order, jobu lapack.Job, jobvt lapack.Job, m int, n int,
	a []float64, lda int, s []float64, u []float64, ldu int,
	vt []float64, ldvt int, superb []float64) lapack.Info {
	info := C.LAPACKE_dgesvd(C.int(order), C.char(jobu), C.char(jobvt),
		C.int(m), C.int(n),
		(*C.double)(unsafe.Pointer(&a[0])), C.int(lda),
		(*C.double)(unsafe.Pointer(&s[0])),
		(*C.double)(unsafe.Pointer(&u[0])), C.int(ldu),
		(*C.double)(unsafe.Pointer(&vt[0])), C.int(ldvt),
		(*C.double)(unsafe.Pointer(&superb[0])))
	return lapack.Info(info)
}

func (Lapack) Dpotrf(order blas.Order, uplo blas.Uplo, n int, a []float64,
	lda int) lapack.Info {
	info := C.LAPACKE_dpotrf(C.int(order), procUplo(uplo), C.int(n),
		(*C.double)(unsafe.Pointer(&a[0])), C.int(lda))
	return lapack.Info(info)
}

func (Lapack) Dsytrf(order blas.Order, uplo blas.Uplo, n int, a []float64,
	lda int, ipiv []int) lapack.Info {
	info := C.LAPACKE_dsytrf(C.int(order), procUplo(uplo), C.int(n),
		(*C.double)(unsafe.Pointer(&a[0])),
		C.int(lda), (*C.int)(unsafe.Pointer(&ipiv[0])))
	return lapack.Info(info)
}
