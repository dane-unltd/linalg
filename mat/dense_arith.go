package mat

import "github.com/gonum/blas"

func (dst *Dense) Add(A Matrix) *Dense {
	ma, na := A.Dims()
	md, nd := dst.Dims()

	if ma != md || na != nd {
		panic("operator dimension mismatch")
	}

	switch A := A.(type) {
	case *Dense:
		dst.Vec(nil).Add(A.Vec(nil))
		return dst
	case *denseView:
		if !A.IsTr() {
			for c := 0; c < na; c++ {
				dst.Col(c, nil).Add(A.Col(c, nil))
			}
			return dst
		}
	}
	for j := 0; j < na; j++ {
		offset := j * dst.rows
		for i := 0; i < ma; i++ {
			dst.data[offset+i] += A.At(i, j)
		}
	}
	return dst
}

func (dst *Dense) AddDiag(d Vec) {
	ma, na := dst.Dims()
	n := na
	if ma < n {
		n = ma
	}
	if len(d) != n {
		panic("dimension mismatch")
	}

	for i, v := range d {
		dst.data[i*dst.rows+i] += v
	}
}

func (dst *Dense) Sub(A Matrix) *Dense {
	ma, na := A.Dims()
	md, nd := dst.Dims()

	if ma != md || na != nd {
		panic("operator dimension mismatch")
	}

	switch A := A.(type) {
	case *Dense:
		dst.Vec(nil).Sub(A.Vec(nil))
		return dst
	case *denseView:
		if !A.IsTr() {
			for c := 0; c < na; c++ {
				dst.Col(c, nil).Sub(A.Col(c, nil))
			}
			return dst
		}
	}
	for j := 0; j < na; j++ {
		offset := j * dst.rows
		for i := 0; i < ma; i++ {
			dst.data[offset+i] -= A.At(i, j)
		}
	}
	return dst
}

func (dst *Dense) ScalCols(v Vec) {
	m, n := dst.Dims()
	if len(v) != n {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		ops.Dscal(m, v[i], dst.Col(i, nil), 1)
	}
}

func (dst *Dense) Mul(A Matrix) *Dense {
	ma, na := A.Dims()
	md, nd := dst.Dims()

	if ma != md || na != nd {
		panic("operator dimension mismatch")
	}

	switch A := A.(type) {
	case *Dense:
		dst.Vec(nil).Mul(A.Vec(nil))
		return dst
	case *denseView:
		if !A.IsTr() {
			for c := 0; c < na; c++ {
				dst.Col(c, nil).Mul(A.Col(c, nil))
			}
			return dst
		}
	}
	for j := 0; j < na; j++ {
		offset := j * dst.rows
		for i := 0; i < ma; i++ {
			dst.data[offset+i] *= A.At(i, j)
		}
	}
	return dst
}

func (dst *Dense) Div(A Matrix) *Dense {
	ma, na := A.Dims()
	md, nd := dst.Dims()

	if ma != md || na != nd {
		panic("operator dimension mismatch")
	}

	switch A := A.(type) {
	case *Dense:
		dst.Vec(nil).Div(A.Vec(nil))
		return dst
	case *denseView:
		if !A.IsTr() {
			for c := 0; c < na; c++ {
				dst.Col(c, nil).Div(A.Col(c, nil))
			}
			return dst
		}
	}
	for j := 0; j < na; j++ {
		offset := j * dst.rows
		for i := 0; i < ma; i++ {
			dst.data[offset+i] /= A.At(i, j)
		}
	}
	return dst
}

func (dst *Dense) MMul(A, B Matrix) {
	ma, na := A.Dims()
	mb, nb := B.Dims()
	m, n := ma, nb

	if na != mb {
		panic("dimension mismatch")
	}

	dst.recvDimCheck(m, n)

	switch A := A.(type) {
	case *Dense:
		switch B := B.(type) {
		case *Dense:
			ops.Dgemm(blas.ColMajor, blas.NoTrans, blas.NoTrans, m, n, na,
				1, A.data, ma, B.data, mb, 0,
				dst.data, m)
			return
		case *denseView:
			ops.Dgemm(blas.ColMajor, blas.NoTrans, B.trans, m, n, na,
				1, A.data, ma, B.data, B.stride, 0,
				dst.data, m)
			return
		}
	case *denseView:
		switch B := B.(type) {
		case *Dense:
			ops.Dgemm(blas.ColMajor, A.trans, blas.NoTrans, m, n, na,
				1, A.data, A.stride, B.data, mb, 0,
				dst.data, m)
			return
		case *denseView:
			ops.Dgemm(blas.ColMajor, A.trans, B.trans, m, n, na,
				1, A.data, A.stride, B.data, B.stride, 0,
				dst.data, m)
			return
		}

	}

	panic("general mmul not implemented")

}

func (dst *Dense) AddSc(a float64) {
	dst.Vec(nil).AddSc(a)
}
