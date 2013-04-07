package lapackinit

import (
	"github.com/dane-unltd/linalg/lapack"
	"github.com/dane-unltd/linalg/matrix"
)

func init() {
	matrix.Ops.Dgesvd = lapack.Dgesvd
	matrix.Ops.Dgeqrf = lapack.Dgeqrf
	matrix.Ops.Dpotrf = lapack.Dpotrf
	matrix.Ops.Dsyevd = lapack.Dsyevd
}
