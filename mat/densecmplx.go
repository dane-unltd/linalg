package mat

import (
	"github.com/gonum/blas"
)

type DenseCmplx struct {
	dense
	data []complex128
}

//Constructors

func NewCmplx(dims ...int) *DenseCmplx {
	if len(dims) == 0 {
		return nil
	}
	m := dims[0]
	n := m
	if len(dims) > 1 {
		n = dims[1]
	}

	D := &DenseCmplx{}
	D.rows = m
	D.cols = n
	D.stride = m
	D.trans = blas.NoTrans

	D.data = make([]complex128, n*m)
	return D
}

func NewFromArrayCmplx(data []complex128, makeCopy bool, dims ...int) *DenseCmplx {
	num := 1
	for _, v := range dims {
		num *= v
	}
	if len(data) < num {
		panic("array length mismatch")
	}

	var D *DenseCmplx
	if !makeCopy {
		if len(dims) == 0 {
			return nil
		}
		m := dims[0]
		n := m
		if len(dims) > 1 {
			n = dims[1]
		}
		D = &DenseCmplx{}
		D.rows = m
		D.cols = n
		D.stride = m
		D.trans = blas.NoTrans
		D.data = data
	} else {
		D = NewCmplx(dims...)
		copy(D.data, data)
	}
	return D
}

//Basic Methods

func reallocCmplx(data *[]complex128, n int, cpy bool) {
	if cap(*data) <= n {
		*data = (*data)[0:n]
		return
	}
	d := make([]complex128, n)
	if cpy {
		copy(d, *data)
	}
	*data = d
}

func zeroCmplx(data []complex128) {
	for i := range data {
		data[i] = 0
	}
}

func (dst *DenseCmplx) recvDimCheck(m, n int) {
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
			dst.data = make([]complex128, m*n)
		} else {
			reallocCmplx(&dst.data, m*n, false)
		}
	}
}
