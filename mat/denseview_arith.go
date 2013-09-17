package mat

import "github.com/gonum/blas"

func (dst *denseView) MMul(A, B Matrix) {
	ma, na := A.Dims()
	mb, nb := B.Dims()
	m, n := ma, nb
	md, nd := dst.Dims()

	if na != mb {
		panic("operand dimension mismatch")
	}
	if md != m || nd != n {
		panic("receiver dimension mismatch")
	}

	if dst.IsTr() {
		panic("cannot use transposed view as receiver")
	}

	switch A := A.(type) {
	case *Dense:
		switch B := B.(type) {
		case *Dense:
			ops.Dgemm(blas.ColMajor, blas.NoTrans, blas.NoTrans, m, n, na,
				1, A.data, ma, B.data, mb, 0,
				dst.data, dst.stride)
			return
		case *denseView:
			ops.Dgemm(blas.ColMajor, blas.NoTrans, B.trans, m, n, na,
				1, A.data, ma, B.data, B.stride, 0,
				dst.data, dst.stride)
			return
		}
	case *denseView:
		switch B := B.(type) {
		case *Dense:
			ops.Dgemm(blas.ColMajor, A.trans, blas.NoTrans, m, n, na,
				1, A.data, A.stride, B.data, mb, 0,
				dst.data, dst.stride)
			return
		case *denseView:
			ops.Dgemm(blas.ColMajor, A.trans, B.trans, m, n, na,
				1, A.data, A.stride, B.data, B.stride, 0,
				dst.data, dst.stride)
			return
		}

	}

	panic("general mmul not implemented")

}

func (D *denseView) ApplyTo(v, dst Vec) {
	m, n := D.Dims()
	if len(v) != n || len(dst) != m {
		panic("dimension mismatch")
	}
	if D.IsTr() {
		m, n = n, m
	}
	ops.Dgemv((blas.ColMajor), D.trans, m, n, 1, D.data, D.stride, v, 1, 0, dst, 1)
}

func (D *denseView) ApplyInverseTo(v, dst Vec) {
	panic("not implemented")
}
