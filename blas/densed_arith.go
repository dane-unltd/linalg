package blas

import "github.com/dane-unltd/linalg/matrix"

type Mulable interface {
	ArrayD() []float64
	Size() (int, int)
	Stride() int
	IsTr() bool
}

func (res DenseD) Mul(A, B matrix.Matrix) {
	matrix.Mul(res, A, B)
}

func (res DenseD) MulDD(A, B *matrix.DenseD) {
	m, n := res.Size()
	_, k := A.Size()
	Dgemm(A.IsTr(), B.IsTr(), m, n, k,
		1, A.ArrayD(), A.Stride(), B.ArrayD(), B.Stride(), 0,
		res.ArrayD(), res.Stride())
}
