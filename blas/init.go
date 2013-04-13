package blas

import "github.com/dane-unltd/linalg/blasops"

func init() {
	blasops.Dasum = Dasum
	blasops.Daxpy = Daxpy
	blasops.Dcopy = Dcopy
	blasops.Ddot = Ddot
	blasops.Dnrm2 = Dnrm2
	blasops.Drot = Drot
	blasops.Drotg = Drotg
	blasops.Dscal = Dscal
	blasops.Dswap = Dswap
	blasops.Idamax = Idamax
	blasops.Isamax = Isamax
	blasops.Sasum = Sasum
	blasops.Saxpy = Saxpy
	blasops.Scopy = Scopy
	blasops.Sdot = Sdot
	blasops.Sdsdot = Sdsdot
	blasops.Snrm2 = Snrm2
	blasops.Srot = Srot
	blasops.Srotg = Srotg
	blasops.Sscal = Sscal
	blasops.Sswap = Sswap

	blasops.Dgemm = Dgemm
}
