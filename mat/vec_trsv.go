package mat

import "github.com/gonum/blas"

func (res Vec) Trsv(A Matrix, b Vec) {
	n := len(res)
	res.Copy(b)
	switch A := A.(type) {
	case *Dense:
		ops.Dtrsv(blas.ColMajor, blas.Upper, blas.NoTrans, blas.NonUnit, n,
			A.data, A.rows, res, 1)
		return
	case *denseView:
		ops.Dtrsv(blas.ColMajor, blas.Upper, A.trans, blas.NonUnit, n,
			A.data, A.stride, res, 1)
		return
	}
	panic("general trsv not implemented")
}
