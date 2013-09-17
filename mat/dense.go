package mat

import "math/rand"
import "github.com/gonum/blas"

type Dense struct {
	rows, cols int
	data       []float64
}

//Constructors

func New(dims ...int) *Dense {
	if len(dims) == 0 {
		return nil
	}
	m := dims[0]
	n := m
	if len(dims) > 1 {
		n = dims[1]
	}

	D := &Dense{}
	D.rows = m
	D.cols = n

	D.data = make([]float64, n*m)
	return D
}

func Xes(x float64, dims ...int) {
	D := New(dims...)
	for i := 0; i < len(D.data); i++ {
		D.data[i] = x
	}
}

func Eye(n int) *Dense {
	D := New(n)
	for i := 0; i < n; i++ {
		D.Set(i, i, 1)
	}
	return D
}

func Rand(dims ...int) *Dense {
	D := New(dims...)
	for i := range D.data {
		D.data[i] = rand.Float64()
	}
	return D
}

func RandN(dims ...int) *Dense {
	D := New(dims...)
	for i := range D.data {
		D.data[i] = rand.NormFloat64()
	}
	return D
}

func NewFromArray(data []float64, makeCopy bool, dims ...int) *Dense {
	num := 1
	for _, v := range dims {
		num *= v
	}
	if len(data) < num {
		panic("array length mismatch")
	}

	var D *Dense
	if !makeCopy {
		if len(dims) == 0 {
			return nil
		}
		m := dims[0]
		n := m
		if len(dims) > 1 {
			n = dims[1]
		}
		D = &Dense{}
		D.rows = m
		D.cols = n
		D.data = data
	} else {
		D = New(dims...)
		copy(D.data, data)
	}
	return D
}

//Basic Methods

func realloc(data *[]float64, n int, cpy bool) {
	if cap(*data) <= n {
		*data = (*data)[0:n]
		return
	}
	d := make([]float64, n)
	if cpy {
		copy(d, *data)
	}
	*data = d
}

func zero(data []float64) {
	for i := range data {
		data[i] = 0
	}
}

func (dst *Dense) recvDimCheck(m, n int) {
	dst.rows = m
	dst.cols = n
	if dst.data == nil {
		dst.data = make([]float64, m*n)
	} else {
		realloc(&dst.data, m*n, false)
	}
}

func (D *Dense) At(i, j int) float64 {
	ix := D.dataIx(i, j)
	return D.data[ix]
}

func (D *Dense) Set(i, j int, v float64) {
	ix := D.dataIx(i, j)
	D.data[ix] = v
}

func (D *Dense) Dims() (int, int) {
	return D.rows, D.cols
}

func (dst *Dense) Copy(A Matrix) *Dense {
	m, n := A.Dims()
	dst.recvDimCheck(m, n)

	switch A := A.(type) {
	case *Dense:
		copy(dst.data, A.data)
	case *denseView:
		if A.trans == blas.NoTrans {
			for c := 0; c < A.cols; c++ {
				copy(dst.data[c*dst.rows:(c+1)*dst.rows], A.data[c*A.stride:])
			}
		}
	}

	for j := 0; j < n; j++ {
		offset := j * dst.rows
		for i := 0; i < m; i++ {
			dst.data[i+offset] = A.At(i, j)
		}
	}
	return dst
}

func (D *Dense) TrView() *denseView {
	return &denseView{
		rows:   D.rows,
		cols:   D.cols,
		stride: D.rows,
		trans:  blas.Trans,
		data:   D.data,
	}
}

func (D *Dense) Equals(A Matrix) bool {
	md, nd := D.Dims()
	ma, na := A.Dims()
	if ma != md || na != nd {
		return false
	}

	switch A := A.(type) {
	case *Dense:
		return Vec(A.data).Equals(Vec(D.data))
	case *denseView:
		if A.IsTr() {
			for i := 0; i < ma; i++ {
				for j := 0; j < na; j++ {
					if A.At(i, j) != D.At(i, j) {
						return false
					}
				}
			}
			return true
		}
		for c := 0; c < A.cols; c++ {
			if !Vec(D.data[c*D.rows : c*D.rows+D.rows]).Equals(Vec(A.data[c*A.stride : c*A.stride+A.rows])) {
				return false
			}
		}
	}
	return true
}

func (D *Dense) Col(ix int, col Vec) Vec {
	m, _ := D.Dims()
	if col == nil {
		return Vec(D.data[ix*D.rows : (ix+1)*D.rows])
	}
	if len(col) != m {
		panic("provided vector has the wrong dimension")
	}
	copy(col, D.data[ix*D.rows:(ix+1)*D.rows])
	return col
}

func (D *Dense) Vec(v Vec) Vec {
	if v == nil {
		return Vec(D.data)
	}
	if len(v) != len(D.data) {
		panic("provided vector has the wrong dimension")
	}
	copy(v, D.data)
	return v
}

func (D *Dense) SetCol(ix int, v Vec) {
	if ix >= D.cols {
		panic("index out of range")
	}
	if v == nil {
		copy(D.data[ix*D.rows:], D.data[(ix+1)*D.rows:])
		D.cols--
		D.data = D.data[:D.cols*D.rows]
		return
	}
	if len(v) != D.rows {
		panic("dimension missmatch")
	}
	copy(D.data[ix*D.rows:], v)
}

func (D *Dense) AddCol(v Vec) {
	D.cols++
	if len(v) != D.rows {
		panic("dimension missmatch")
	}

	realloc(&D.data, D.cols*D.rows, true)
	copy(D.data[(D.cols-1)*D.rows:], v)
}

func (D *Dense) View(i, j, r, c int) *denseView {
	m, n := D.Dims()
	if i < 0 || j < 0 {
		panic("negative index")
	}
	if i+r > m || j+c > n {
		panic("submatrix out of bounds")
	}

	ix := D.dataIx(i, j)
	return &denseView{
		rows:   r,
		cols:   c,
		stride: D.rows,
		trans:  blas.NoTrans,
		data:   D.data[ix : ix+(c-1)*D.rows+r],
	}
}

func (dst *Dense) SubMatrix(A Matrix, i, j, r, c int) {
	if A, ok := A.(Viewer); ok {
		V := A.View(i, j, r, c)
		dst.Copy(V)
		return
	}

	m, n := A.Dims()
	if i < 0 || j < 0 {
		panic("negative index")
	}
	if i+r > m || j+c > n {
		panic("submatrix out of bounds")
	}

	dst.recvDimCheck(r, c)
	for colIx := 0; colIx < c; colIx++ {
		for rowIx := 0; rowIx < r; rowIx++ {
			dst.data[rowIx*r+colIx] = A.At(rowIx, colIx)
		}
	}
}

func (D *Dense) dataIx(i, j int) int {
	return i + j*D.rows
}
