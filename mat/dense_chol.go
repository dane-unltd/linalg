package mat

import "github.com/dane-unltd/linalg/lapack"

func (D *Dense) Chol(R *Dense) lapack.Info {
	m, n := D.Size()
	if m != n {
		panic("need square matrix for chol")
	}
	ops.Dlacpy('U', m, n, D.data, D.stride, R.data, R.stride)
	return ops.Dpotrf('U', m, R.data, R.stride)
}
