package mat

import "fmt"

import "github.com/gonum/blas"

func (dst *Dense) Add(A, B *Dense) {
	ma, na := A.Dims()
	mb, nb := B.Dims()

	if ma != mb || na != nb {
		fmt.Println(ma, na, mb, nb)
		panic("operator dimension mismatch")
	}

	dst.recvDimCheck(ma, na)

	if (!A.IsTr() && !B.IsTr() && !dst.IsTr()) ||
		(A.IsTr() && B.IsTr() && dst.IsTr()) {
		for c := 0; c < na; c++ {
			dst.ColView(c).Add(A.ColView(c), B.ColView(c))
		}
		return
	}
	for i := 0; i < ma; i++ {
		for j := 0; j < na; j++ {
			dst.Set(i, j, A.At(i, j)+B.At(i, j))
		}
	}
}

func (dst *Dense) AddDiag(A *Dense, d Vec) {
	ma, na := A.Dims()
	n := na
	if ma < n {
		n = ma
	}
	if len(d) != n {
		panic("dimension mismatch")
	}

	dst.Copy(A)
	for i := range d {
		dst.Set(i, i, dst.At(i, i)+d[i])
	}
}

func (dst *Dense) Sub(A, B *Dense) {
	ma, na := A.Dims()
	mb, nb := B.Dims()

	if ma != mb || na != nb {
		fmt.Println(ma, na, mb, nb)
		panic("operator dimension mismatch")
	}

	dst.recvDimCheck(ma, na)

	if (!A.IsTr() && !B.IsTr() && !dst.IsTr()) ||
		(A.IsTr() && B.IsTr() && dst.IsTr()) {
		for c := 0; c < na; c++ {
			dst.ColView(c).Sub(A.ColView(c), B.ColView(c))
		}
		return
	}
	for i := 0; i < ma; i++ {
		for j := 0; j < na; j++ {
			dst.Set(i, j, A.At(i, j)-B.At(i, j))
		}
	}
}

func (dst *Dense) ScalCols(A *Dense, v Vec) {
	m, n := A.Dims()
	if len(v) != n {
		panic("dimension mismatch")
	}
	dst.recvDimCheck(m, n)

	if dst.IsTr() || A.IsTr() {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				dst.Set(i, j, A.At(i, j)*v[j])
			}
		}
		return
	}

	for i := 0; i < n; i++ {
		ops.Dcopy(m, A.ColView(i), 1, dst.ColView(i), 1)
		ops.Dscal(m, v[i], dst.ColView(i), 1)
	}
}

func (D *Dense) ApplyTo(v, dst Vec) {
	m, n := D.Dims()
	if len(v) != n || len(dst) != m {
		panic("dimension mismatch")
	}
	if D.IsTr() {
		m, n = n, m
	}
	ops.Dgemv((blas.ColMajor), (D.trans), m, n, 1, D.data, D.stride, v, 1, 0, dst, 1)
}

func (dst *Dense) MulElem(A, B *Dense) {
	ma, na := A.Dims()
	mb, nb := B.Dims()

	if ma != mb || na != nb {
		fmt.Println(ma, na, mb, nb)
		panic("operator dimension mismatch")
	}

	dst.recvDimCheck(ma, na)

	if (!A.IsTr() && !B.IsTr() && !dst.IsTr()) ||
		(A.IsTr() && B.IsTr() && dst.IsTr()) {
		for c := 0; c < na; c++ {
			dst.ColView(c).Mul(A.ColView(c), B.ColView(c))
		}
		return
	}

	for i := 0; i < ma; i++ {
		for j := 0; j < na; j++ {
			dst.Set(i, j, A.At(i, j)*B.At(i, j))
		}
	}
}

func (dst *Dense) Mul(A, B *Dense) {
	ma, na := A.Dims()
	mb, nb := B.Dims()
	m, n := ma, nb

	if na != mb {
		fmt.Println(m, n, ma, na, mb, nb)
		panic("dimension mismatch")
	}

	dst.recvDimCheck(m, n)

	if dst.IsTr() {
		A.transp()
		B.transp()
		ops.Dgemm(blas.ColMajor, B.trans, A.trans, n, m, na,
			1, B.data, B.stride, A.data, A.stride, 0,
			dst.data, dst.stride)
		A.transp()
		B.transp()
		return
	}

	ops.Dgemm(blas.ColMajor, A.trans, B.trans, m, n, na,
		1, A.data, A.stride, B.data, B.stride, 0,
		dst.data, dst.stride)
}

func (dst *Dense) AddSc(a float64) {
	for j := 0; j < dst.cols; j++ {
		for i := 0; i < dst.rows; i++ {
			dst.data[j*dst.stride+i] += a
		}
	}
}
