package goblas

import "github.com/dane-unltd/linalg/blas"

func init() {
	blas.Dasum = Dasum
	blas.Daxpy = Daxpy
	blas.Dcopy = Dcopy
	blas.Ddot = Ddot
	blas.Dnrm2 = Dnrm2
	blas.Drot = Drot
	blas.Drotg = Drotg
	blas.Dscal = Dscal
	blas.Dswap = Dswap
	blas.Idamax = Idamax
	blas.Isamax = Isamax
	blas.Sasum = Sasum
	blas.Saxpy = Saxpy
	blas.Scopy = Scopy
	blas.Sdot = Sdot
	blas.Sdsdot = Sdsdot
	blas.Snrm2 = Snrm2
	blas.Srot = Srot
	blas.Srotg = Srotg
	blas.Sscal = Sscal
	blas.Sswap = Sswap

	blas.Dgemm = Dgemm
}
