package mat

import "github.com/gonum/blas"

type denseView struct {
	rows, cols int
	stride     int
	trans      blas.Transpose
	data       []float64
}

func (D *denseView) At(i, j int) float64 {
	return D.data[D.dataIx(i, j)]
}

func (D *denseView) Dims() (int, int) {
	if D.IsTr() {
		return D.cols, D.rows
	}
	return D.rows, D.cols
}

func (D *denseView) IsTr() bool {
	return D.trans == blas.Trans
}

func (D *denseView) View(i, j, r, c int) *denseView {
	m, n := D.Dims()
	if i < 0 || j < 0 {
		panic("negative index")
	}
	if i+r > m || j+c > n {
		panic("submatrix out of bounds")
	}
	A := *D

	if A.IsTr() {
		A.rows = c
		A.cols = r
	} else {
		A.rows = r
		A.cols = c
	}
	ix := A.dataIx(i, j)
	A.data = A.data[ix : ix+(A.cols-1)*A.stride+A.rows]
	return &A
}

func (D *denseView) TrView() *denseView {
	Dt := *D
	if Dt.IsTr() {
		Dt.trans = blas.NoTrans
	} else {
		Dt.trans = blas.Trans
	}
	return &Dt
}

func (D *denseView) dataIx(i, j int) int {
	if D.IsTr() {
		return j + i*D.stride
	}
	return i + j*D.stride
}

func (D *denseView) Col(ix int, col Vec) Vec {
	m, _ := D.Dims()
	if col == nil {
		if D.IsTr() {
			panic("cannot take col view of transposed view")
		}
		return Vec(D.data[ix*D.stride : ix*D.stride+D.rows])
	}
	if len(col) != m {
		panic("provided vector has the wrong dimension")
	}

	if D.IsTr() {
		for row := 0; row < m; row++ {
			col[row] = D.data[ix+row*D.stride]
		}
	} else {
		copy(col, D.data[ix*D.rows:(ix+1)*D.rows])
	}
	return col
}

func (dst *denseView) Copy(A Matrix) *denseView {
	if dst.IsTr() {
		panic("cannot use transposed view as receiver")
	}

	ma, na := A.Dims()
	m, n := dst.Dims()

	if ma != m || na != n {
		panic("dimension mismatch")
	}

	switch A := A.(type) {
	case *Dense:
		for c := 0; c < A.cols; c++ {
			copy(dst.data[c*dst.stride:c*dst.stride+dst.rows],
				A.data[c*A.rows:])
		}
	case *denseView:
		if A.trans == blas.NoTrans {
			for c := 0; c < A.cols; c++ {
				copy(dst.data[c*dst.stride:c*dst.stride+dst.rows],
					A.data[c*A.stride:])
			}
		}
	}

	for j := 0; j < n; j++ {
		offset := j * dst.stride
		for i := 0; i < m; i++ {
			dst.data[i+offset] = A.At(i, j)
		}
	}
	return dst
}
