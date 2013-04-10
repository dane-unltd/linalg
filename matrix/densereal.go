package matrix

import "math/rand"

type ScD float64

type MatDable interface {
	Size() (int, int)
	Stride() int

	ArrayD() []float64
	IsTr() bool
}

type MatD struct {
	dense
	data []float64
}

func (d ScD) Size() (int, int) {
	return 1, 1
}

func (d ScD) Strides() (int, int) {
	return 1, 1
}

func (d ScD) ArrayD() []float64 {
	return []float64{float64(d)}
}

func (d ScD) At(ixs ...int) float64 {
	return float64(d)
}

func Zeros(dims ...int) *MatD {
	if len(dims) == 0 {
		return nil
	}
	m := dims[0]
	n := m
	if len(dims) > 1 {
		n = dims[1]
	}

	D := &MatD{}
	D.rows = m
	D.cols = n
	D.stride = m

	D.data = make([]float64, n*m)
	return D
}

func Eye(n int) *MatD {
	D := Zeros(n)
	for i := 0; i < n; i++ {
		D.Set(1, i, i)
	}
	return D
}

func Diag(d []float64) *MatD {
	D := Zeros(len(d))
	for i := 0; i < len(d); i++ {
		D.Set(d[i], i, i)
	}
	return D
}

func RandU(dims ...int) *MatD {
	D := Zeros(dims...)
	for i := range D.data {
		D.data[i] = rand.Float64()
	}
	return D
}

func RandN(dims ...int) *MatD {
	D := Zeros(dims...)
	for i := range D.data {
		D.data[i] = rand.NormFloat64()
	}
	return D
}

func FromArrayD(data []float64, useArray bool, dims ...int) *MatD {
	num := 1
	for _, v := range dims {
		num *= v
	}
	if len(data) < num {
		panic("array length mismatch")
	}

	var D *MatD
	if useArray {
		if len(dims) == 0 {
			return nil
		}
		m := dims[0]
		n := m
		if len(dims) > 1 {
			n = dims[1]
		}
		D = &MatD{}
		D.rows = m
		D.cols = n
		D.stride = m
		D.data = data
	} else {
		D = Zeros(dims...)
		copy(D.data, data)
	}
	return D
}

func (D *MatD) Copy() *MatD {
	Dc := MatD{}
	Dc.rows = D.rows
	Dc.cols = D.cols
	Dc.stride = D.stride
	Dc.trans = D.trans
	Dc.data = make([]float64, len(D.data))
	copy(Dc.data, D.data)
	return &Dc
}

func (D *MatD) Set(v float64, ixs ...int) {
	ix := -1
	if len(ixs) == 1 {
		ix = D.dataIx(ixs[0])
	} else {
		if D.trans {
			ix = ixs[1] + ixs[0]*D.stride
		} else {
			ix = ixs[0] + ixs[1]*D.stride
		}
	}
	D.data[ix] = v
}

func (D *MatD) Col(ix int) VecD {
	return VecD(D.data[ix*D.stride : (ix*D.stride + D.rows)])
}

func (D *MatD) SetCol(ix int, v VecD) {
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

func (D *MatD) AddCol(v VecD) {
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

func (D *MatD) At(ixs ...int) float64 {
	ix := -1
	if len(ixs) == 1 {
		ix = D.dataIx(ixs[0])
	} else {
		if D.trans {
			ix = ixs[1] + ixs[0]*D.stride
		} else {
			ix = ixs[0] + ixs[1]*D.stride
		}
	}
	return D.data[ix]
}

func (D *MatD) ArrayD() []float64 {
	return D.data
}

func (D *MatD) dataIx(matIx int) int {
	n := matIx / D.rows
	m := matIx - n*D.rows
	return m + n*D.stride
}
