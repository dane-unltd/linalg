package mat

import "github.com/dane-unltd/linalg/lapack"
import "github.com/gonum/blas"

func (D *Dense) Chol(R *Dense) lapack.Info {
	m, n := D.Dims()
	if m != n {
		panic("need square matrix for chol")
	}
	R.recvDimCheck(m, n)
	if R.IsTr() {
		panic("cannot store factorization into transposed view")
	}
	for i := 0; i < n; i++ {
		copy(R.data[i*R.stride:i*R.stride+i+1], D.data[i*D.stride:])
	}
	return ops.Dpotrf(blas.ColMajor, blas.Upper, m, R.data, R.stride)
}
