package matrix

import "math/rand"
import "github.com/dane-unltd/linalg/blas"

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

type DenseFloat64 struct {
	dense
	data []float64
}

func NewDenseFloat64(dims ...int) *DenseFloat64 {
	if len(dims) == 0 {
		return nil
	}
	m := dims[0]
	n := m
	if len(dims) > 1 {
		n = dims[1]
	}

	D := &DenseFloat64{}
	D.rows = m
	D.cols = n
	D.stride = m
	D.trans = blas.NoTrans

	D.data = make([]float64, n*m)
	return D
}

func Eye(n int) *DenseFloat64 {
	D := NewDenseFloat64(n)
	for i := 0; i < n; i++ {
		D.Set(i, i, 1)
	}
	return D
}

func RandU(dims ...int) *DenseFloat64 {
	D := NewDenseFloat64(dims...)
	for i := range D.data {
		D.data[i] = rand.Float64()
	}
	return D
}

func RandN(dims ...int) *DenseFloat64 {
	D := NewDenseFloat64(dims...)
	for i := range D.data {
		D.data[i] = rand.NormFloat64()
	}
	return D
}

func FromArrayD(data []float64, useArray bool, dims ...int) *DenseFloat64 {
	num := 1
	for _, v := range dims {
		num *= v
	}
	if len(data) < num {
		panic("array length mismatch")
	}

	var D *DenseFloat64
	if useArray {
		if len(dims) == 0 {
			return nil
		}
		m := dims[0]
		n := m
		if len(dims) > 1 {
			n = dims[1]
		}
		D = &DenseFloat64{}
		D.rows = m
		D.cols = n
		D.stride = m
		D.trans = blas.NoTrans
		D.data = data
	} else {
		D = NewDenseFloat64(dims...)
		copy(D.data, data)
	}
	return D
}

func (D *DenseFloat64) Copy() interface{} {
	Dc := *D
	Dc.data = make([]float64, len(D.data))
	copy(Dc.data, D.data)
	return &Dc
}

func (D *DenseFloat64) Equals(x interface{}) bool {
	D2, ok := x.(*DenseFloat64)
	if !ok {
		return ok
	}
	return D.VecView().Equals(D2.VecView())
}

func (D *DenseFloat64) At(i, j int) float64 {
	ix := D.dataIx(i, j)
	return D.data[ix]
}

func (D *DenseFloat64) Set(i, j int, v float64) {
	ix := D.dataIx(i, j)
	D.data[ix] = v
}

func (D *DenseFloat64) T() *DenseFloat64 {
	if D.IsTr() {
		D.trans = blas.NoTrans
		return D
	}

	D.trans = blas.Trans

	return D
}

func (D *DenseFloat64) ColView(ix int) VecFloat64 {
	return VecFloat64(D.data[ix*D.stride : (ix*D.stride + D.rows)])
}

func (D *DenseFloat64) VecView() VecFloat64 {
	return VecFloat64(D.data)
}

func (D *DenseFloat64) SetCol(ix int, v VecFloat64) {
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

func (D *DenseFloat64) AddCol(v VecFloat64) {
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

func (D *DenseFloat64) Array() []float64 {
	return D.data
}

func (D *DenseFloat64) dataIx(i, j int) int {
	if D.IsTr() {
		return j + i*D.stride
	}
	return i + j*D.stride
}
