package matrix

import "math/rand"
import blas "github.com/dane-unltd/linalg/blasops"

type DenseD struct {
	dense
	data []float64
}

func NewDenseD(dims ...int) *DenseD {
	if len(dims) == 0 {
		return nil
	}
	m := dims[0]
	n := m
	if len(dims) > 1 {
		n = dims[1]
	}

	D := &DenseD{}
	D.rows = m
	D.cols = n
	D.stride = m
	D.trans = blas.NoTrans

	D.data = make([]float64, n*m)
	return D
}

func Eye(n int) *DenseD {
	D := NewDenseD(n)
	for i := 0; i < n; i++ {
		D.Set(i, i, 1)
	}
	return D
}

func RandU(dims ...int) *DenseD {
	D := NewDenseD(dims...)
	for i := range D.data {
		D.data[i] = rand.Float64()
	}
	return D
}

func RandN(dims ...int) *DenseD {
	D := NewDenseD(dims...)
	for i := range D.data {
		D.data[i] = rand.NormFloat64()
	}
	return D
}

func FromArrayD(data []float64, useArray bool, dims ...int) *DenseD {
	num := 1
	for _, v := range dims {
		num *= v
	}
	if len(data) < num {
		panic("array length mismatch")
	}

	var D *DenseD
	if useArray {
		if len(dims) == 0 {
			return nil
		}
		m := dims[0]
		n := m
		if len(dims) > 1 {
			n = dims[1]
		}
		D = &DenseD{}
		D.rows = m
		D.cols = n
		D.stride = m
		D.trans = blas.NoTrans
		D.data = data
	} else {
		D = NewDenseD(dims...)
		copy(D.data, data)
	}
	return D
}

func (D *DenseD) Copy() interface{} {
	Dc := *D
	Dc.data = make([]float64, len(D.data))
	copy(Dc.data, D.data)
	return &Dc
}

func (D *DenseD) Equals(x interface{}) bool {
	D2, ok := x.(*DenseD)
	if !ok {
		return ok
	}
	return D.VecView().Equals(D2.VecView())
}

func (D *DenseD) At(i, j int) float64 {
	ix := D.dataIx(i, j)
	return D.data[ix]
}

func (D *DenseD) Set(i, j int, v float64) {
	ix := D.dataIx(i, j)
	D.data[ix] = v
}

func (D *DenseD) T() *DenseD {
	if D.IsTr() {
		D.trans = blas.NoTrans
		return D
	}

	D.trans = blas.Trans

	return D
}

func (D *DenseD) ColView(ix int) VecD {
	return VecD(D.data[ix*D.stride : (ix*D.stride + D.rows)])
}

func (D *DenseD) VecView() VecD {
	return VecD(D.data)
}

func (D *DenseD) SetCol(ix int, v VecD) {
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

func (D *DenseD) AddCol(v VecD) {
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

func (D *DenseD) ArrayD() []float64 {
	return D.data
}

func (D *DenseD) dataIx(i, j int) int {
	if D.IsTr() {
		return j + i*D.stride
	}
	return i + j*D.stride
}
