package matrix

import "github.com/gonum/blas"

func (res Vec) Trsv(A *Dense, b Vec) {
	n := len(res)
	res.Copy(b)
	ops.Dtrsv(blas.ColMajor, 'U', A.trans, blas.NonUnit, n, A.data, A.stride, res, 1)
}
