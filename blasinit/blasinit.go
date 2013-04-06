package blasinit

import "github.com/dane-unltd/linalg/matrix"
import "github.com/dane-unltd/linalg/blas"

func init() {
	matrix.Ops.Dnrm2 = blas.Dnrm2
	matrix.Ops.Dasum = blas.Dasum
	matrix.Ops.Ddot = blas.Ddot
	matrix.Ops.Idamax = blas.Idamax
	matrix.Ops.Dswap = blas.Dswap
	matrix.Ops.Dcopy = blas.Dcopy
	matrix.Ops.Daxpy = blas.Daxpy
	matrix.Ops.Dscal = blas.Dscal
	matrix.Ops.Dgemm = blas.Dgemm
}
