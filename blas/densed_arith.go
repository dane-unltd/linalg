package blas

import "github.com/dane-unltd/linalg/matrix"
import "reflect"

//registering fast paths
func init() {
	tp := matrix.TypePair{
		reflect.TypeOf((*matrix.DenseD)(nil)),
		reflect.TypeOf((*matrix.DenseD)(nil)),
	}
	matrix.DenseDMul[tp] = DenseDMulAA
}

func DenseDMulAA(res *matrix.DenseD, Am, Bm matrix.Matrix) {
	A := Am.(*matrix.DenseD)
	B := Bm.(*matrix.DenseD)

	m, n := res.Size()
	_, k := A.Size()
	Dgemm(A.IsTr(), B.IsTr(), m, n, k,
		1, A.ArrayD(), A.Stride(), B.ArrayD(), B.Stride(), 0,
		res.ArrayD(), res.Stride())
}
