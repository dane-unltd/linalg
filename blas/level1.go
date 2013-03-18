package blas

import (
	//"errors"
	"github.com/dane-unltd/linalg/dense"
	"math"
)

// Returns the Euclidean norm of a vector (returns ||x||_2). 
func Nrm2D(X MatD) float64 {
	m, n := X.Size()
	stride := X.Stride()
	data, _ := X.ArrayD()
	if n == 1 {
		return dnrm2(m, data, 1)
	}
	if m == 1 {
		return dnrm2(n, data, stride)
	}
	v := 0.0
	for i := 0; i < n; i++ {
		v += dnrm2(m, data[i*stride:], 1)
	}
	return v
}

// Returns ||Re x||_1 + ||Im x||_1.
func AsumD(X MatD) float64 {
	m, n := X.Size()
	stride := X.Stride()
	data, _ := X.ArrayD()
	if n == 1 {
		return dasum(m, data, 1)
	}
	if m == 1 {
		return dasum(n, data, stride)
	}
	v := 0.0
	for i := 0; i < n; i++ {
		v += dasum(m, data[i*stride:], 1)
	}
	return v
}

// Returns Y = X^T*Y
func DotD(X, Y MatD) float64 {
	mX, nX := X.Size()
	mY, nY := Y.Size()
	if nX != 1 || nY != 1 || mX != mY {
		return math.NaN()
	}
	Xa, _ := X.ArrayD()
	Ya, _ := Y.ArrayD()
	return ddot(mX, Xa, 1, Ya, 1)
}

func Max(X MatD, dim int) (*dense.MatD, []int) {
	m, n := X.Size()
	stride := X.Stride()
	data, _ := X.ArrayD()

	var res *dense.MatD
	var ixs []int
	if dim == 1 {
		res = dense.Zeros(1, n)
		ixs = make([]int, n)
		for i := range ixs {
			ixs[i] = idamax(n, data[i:], stride) - 1
			res.Set(data[i+ixs[i]*stride], i)
		}
	} else if dim == 0 {
		res = dense.Zeros(m, 1)
		ixs = make([]int, m)
		for i := range ixs {
			ixs[i] = idamax(m, data[i*stride:], 1) - 1
			res.Set(data[i*stride+ixs[i]], i)
		}
	} else {
		panic("dim must be 0 or 1")
	}
	return res, ixs
}
