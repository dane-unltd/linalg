package matrix

type Muler interface {
	MulMM(A, B Matrix)
	Size() (int, int)
	Copy() Matrix
}

type MulDDer interface {
	MulDD(A, B *DenseD)
}

type MulMVecer interface {
	MulMVec(A Matrix, v VecD)
}

type MulMDiager interface {
	MulMDiag(A Matrix, D DiagD)
}

func Mul(res Muler, A, B Matrix) {
	m, n := res.Size()
	ma, na := A.Size()
	mb, nb := B.Size()

	if ma != m || nb != n || na != mb {
		panic("dimension missmatch")
	}

	switch A := A.(type) {
	case *DenseD:
		switch B := B.(type) {
		case *DenseD:
			if res, ok := res.(MulDDer); ok {
				res.MulDD(A, B)
				return
			}
		}
	default:
		switch B := B.(type) {
		case DiagD:
			if res, ok := res.(MulMDiager); ok {
				res.MulMDiag(A, B)
				return
			}
		case VecD:
			if res, ok := res.(MulMVecer); ok {
				res.MulMVec(A, B)
				return
			}
		}
	}
	res.MulMM(A, B)
}
