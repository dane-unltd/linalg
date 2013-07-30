package mat

import "math/rand"
import "github.com/gonum/blas"

type dense struct {
	rows, cols int
	stride     int
	trans      blas.Transpose
	view       bool
}

type Dense struct {
	dense
	data []float64
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
	D.stride = m
	D.trans = blas.NoTrans

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
		D.stride = m
		D.trans = blas.NoTrans
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
	md, nd := dst.Dims()
	if dst.view {
		if m != md || n != nd {
			panic("receiver dimension missmatch")
		}
	} else {
		dst.rows = m
		dst.cols = n
		dst.stride = m
		dst.trans = blas.NoTrans
		if dst.data == nil {
			dst.data = make([]float64, m*n)
		} else {
			realloc(&dst.data, m*n, false)
		}
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

func (D *dense) Dims() (int, int) {
	if D.IsTr() {
		return D.cols, D.rows
	}
	return D.rows, D.cols
}

func (D *dense) Stride() int {
	return D.stride
}

func (D *dense) IsTr() bool {
	return D.trans == blas.Trans
}

func (dst *Dense) Copy(A *Dense) {
	m, n := A.Dims()

	dst.recvDimCheck(m, n)

	if A.trans != dst.trans {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				dst.Set(i, j, A.At(i, j))
			}
		}
		return
	}

	if A.stride == m && dst.stride == m {
		copy(dst.data, A.data)
	} else {
		for c := 0; c < A.cols; c++ {
			copy(dst.data[c*dst.stride:(c+1)*dst.stride], A.data[c*A.stride:])
		}
	}

}

func (D *Dense) TrView() *Dense {
	Dt := *D
	Dt.transp()
	Dt.view = true
	return &Dt
}

func (D *Dense) Equals(A *Dense) bool {
	md, nd := D.Dims()
	ma, na := A.Dims()
	if ma != md || na != nd {
		return false
	}

	if A.trans != D.trans {
		for i := 0; i < ma; i++ {
			for j := 0; j < na; j++ {
				if A.At(i, j) != D.At(i, j) {
					return false
				}
			}
		}
		return true
	}

	if A.stride == ma && D.stride == md {
		return Vec(A.data).Equals(Vec(D.data))
	} else {
		for c := 0; c < A.cols; c++ {
			if !Vec(D.data[c*D.stride : c*D.stride+D.rows]).Equals(Vec(A.data[c*A.stride : c*A.stride+A.rows])) {
				return false
			}
		}
	}
	return true
}

func (D *Dense) transp() *Dense {
	if D.IsTr() {
		D.trans = blas.NoTrans
		return D
	}

	D.trans = blas.Trans

	return D
}

func (D *Dense) Col(ix int, col Vec) Vec {
	m, _ := D.Dims()
	if col == nil && !D.IsTr() {
		return D.ColView(ix)
	}
	if col == nil {
		col = make(Vec, m)
	}
	if len(col) != m {
		panic("provided vector has the wrong dimension")
	}

	if !D.IsTr() {
		copy(col, D.data[ix*D.stride:(ix*D.stride+D.rows)])
	} else {
		for i := 0; i < m; i++ {
			col[i] = D.At(i, ix)
		}
	}
	return col
}

func (D *Dense) ColView(ix int) Vec {
	if D.IsTr() {
		panic("cannot get column view of transposed matrix")
	}
	return Vec(D.data[ix*D.stride : (ix*D.stride + D.rows)])
}

func (D *Dense) VecView() Vec {
	if D.view {
		panic("cannot get vectorized view of a view")
	}
	return Vec(D.data)
}

func (D *Dense) SetCol(ix int, v Vec) {
	if D.IsTr() {
		panic("column operation on transposed matrix not implemented")
	}
	if ix >= D.cols {
		panic("index out of range")
	}
	if v == nil {
		copy(D.data[ix*D.stride:], D.data[(ix+1)*D.stride:])
		D.cols--
		D.data = D.data[:D.cols*D.stride]
		return
	}
	if len(v) != D.rows {
		panic("dimension missmatch")
	}
	copy(D.data[ix*D.stride:], v)
}

func (D *Dense) AddCol(v Vec) {
	D.cols++
	if len(v) != D.rows {
		panic("dimension missmatch")
	}

	realloc(&D.data, D.cols*D.stride, true)
	copy(D.data[(D.cols-1)*D.stride:], v)
}

func (D *Dense) View(i, j, r, c int) *Dense {
	m, n := D.Dims()
	if i < 0 || j < 0 {
		panic("negative index")
	}
	if i+r > m || j+c > n {
		panic("submatrix out of bounds")
	}
	A := *D
	A.view = true

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

func (dst *Dense) SubMatrix(A *Dense, i, j, r, c int) {
	V := A.View(i, j, r, c)
	dst.Copy(V)
}

func (D *Dense) Array() []float64 {
	return D.data
}

func (D *Dense) dataIx(i, j int) int {
	if D.IsTr() {
		return j + i*D.stride
	}
	return i + j*D.stride
}
