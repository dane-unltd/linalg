package lapackinit

import (
	"github.com/dane-unltd/linalg/lapack"
	"github.com/dane-unltd/linalg/matrix"
)

func init() {
	matrix.Ops.Dgesvd = lapack.Dgesvd
}
