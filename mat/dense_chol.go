package mat

import "github.com/dane-unltd/linalg/lapack"
import "github.com/gonum/blas"

func (D *Dense) Chol(R *Dense) lapack.Info {
	m, n := D.Dims()
	if m != n {
		panic("need square matrix for chol")
	}
	R.recvDimCheck(m, n)

	for i := 0; i < n; i++ {
		copy(R.data[i*R.rows:i*R.rows+i+1], D.data[i*D.rows:])
	}
	return ops.Dpotrf(blas.ColMajor, blas.Upper, m, R.data, R.rows)
}
