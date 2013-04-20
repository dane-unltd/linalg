package matrix

import "math/rand"
import "github.com/kortschak/blas"

type dense struct {
	rows, cols int
	stride     int
	trans      blas.Transpose
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

func RandU(dims ...int) *Dense {
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

func FromArray(data []float64, useArray bool, dims ...int) *Dense {
	num := 1
	for _, v := range dims {
		num *= v
	}
	if len(data) < num {
		panic("array length mismatch")
	}

	var D *Dense
	if useArray {
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

func (D *Dense) Copy() interface{} {
	Dc := *D
	Dc.data = make([]float64, len(D.data))
	copy(Dc.data, D.data)
	return &Dc
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

func (D *Dense) T() *Dense {
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
