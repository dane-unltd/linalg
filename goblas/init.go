package goblas

import "github.com/dane-unltd/linalg/blas"

func init() {
	ops := blas.BlasOps{}
	ops.Dasum = Dasum
	ops.Daxpy = Daxpy
	ops.Dcopy = Dcopy
	ops.Ddot = Ddot
	ops.Dnrm2 = Dnrm2
	ops.Drot = Drot
	ops.Drotg = Drotg
	ops.Dscal = Dscal
	ops.Dswap = Dswap
	ops.Idamax = Idamax
	ops.Isamax = Isamax
	ops.Sasum = Sasum
	ops.Saxpy = Saxpy
	ops.Scopy = Scopy
	ops.Sdot = Sdot
	ops.Sdsdot = Sdsdot
	ops.Snrm2 = Snrm2
	ops.Srot = Srot
	ops.Srotg = Srotg
	ops.Sscal = Sscal
	ops.Sswap = Sswap

	ops.Dgemm = Dgemm

	blas.Register(ops)
}
