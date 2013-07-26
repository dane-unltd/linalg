package mat

import "math/rand"
import "github.com/gonum/blas"

type dense struct {
	rows, cols int
	stride     int
	trans      blas.Transpose
	view       bool
}

func (D *dense) Size() (int, int) {
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

type Dense struct {
	dense
	data []float64
}

func NewDense(dims ...int) *Dense {
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

func Eye(n int) *Dense {
	D := NewDense(n)
	for i := 0; i < n; i++ {
		D.Set(i, i, 1)
	}
	return D
}

func Rand(dims ...int) *Dense {
	D := NewDense(dims...)
	for i := range D.data {
		D.data[i] = rand.Float64()
	}
	return D
}

func RandN(dims ...int) *Dense {
	D := NewDense(dims...)
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
		D = NewDense(dims...)
		copy(D.data, data)
	}
	return D
}

func (D *Dense) Copy(A Matrix) {
	m, n := A.Size()
	if len(D.data) < m*n {
		panic("not enough space for copy")
	}
	switch A := A.(type) {
	case *Dense:
		D.dense = A.dense
		D.stride = A.rows
		if A.stride == A.rows {
			copy(D.data[:m*n], A.data)
		} else {
			for c := 0; c < A.cols; c++ {
				copy(D.data[c*D.stride:(c+1)*D.stride], A.data[c*A.stride:])
			}
		}
		return
	}

	D.rows = m
	D.stride = m
	D.cols = n
	D.trans = blas.NoTrans

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			D.Set(i, j, A.At(i, j))
		}
	}
}

func (D *Dense) TrView() *Dense {
	Dt := *D
	Dt.transp()
	Dt.view = true
	return &Dt
}

func (D *Dense) Equals(x interface{}) bool {
	D2, ok := x.(*Dense)
	if !ok {
		return ok
	}
	return D.VecView().Equals(D2.VecView())
}

func (D *Dense) At(i, j int) float64 {
	ix := D.dataIx(i, j)
	return D.data[ix]
}

func (D *Dense) Set(i, j int, v float64) {
	ix := D.dataIx(i, j)
	D.data[ix] = v
}

func (D *Dense) transp() *Dense {
	if D.IsTr() {
		D.trans = blas.NoTrans
		return D
	}

	D.trans = blas.Trans

	return D
}

func (D *Dense) ColView(ix int) Vec {
	return Vec(D.data[ix*D.stride : (ix*D.stride + D.rows)])
}

func (D *Dense) VecView() Vec {
	return Vec(D.data)
}

func (D *Dense) SetCol(ix int, v Vec) {
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

	if len(D.data) < D.cols*D.stride {
		data := make([]float64, D.cols*D.stride)
		copy(data, D.data)
		copy(data[(D.cols-1)*D.stride:], v)
		D.data = data
	} else {
		copy(D.data[(D.cols-1)*D.stride:], v)
	}
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
