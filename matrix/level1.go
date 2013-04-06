package matrix

import (
	//"errors"
	"math"
)

// Returns the Euclidean norm of a vector (returns ||x||_2). 
func Nrm2D(X MatDable) float64 {
	m, n := X.Size()
	stride := X.Stride()
	data := X.ArrayD()
	if n == 1 {
		return Ops.Dnrm2(m, data, 1)
	}
	if m == 1 {
		return Ops.Dnrm2(n, data, stride)
	}
	v := 0.0
	for i := 0; i < n; i++ {
		v += Ops.Dnrm2(m, data[i*stride:], 1)
	}
	return v
}

// Returns ||Re x||_1 + ||Im x||_1.
func AsumD(X MatDable) float64 {
	m, n := X.Size()
	stride := X.Stride()
	data := X.ArrayD()
	if n == 1 {
		return Ops.Dasum(m, data, 1)
	}
	if m == 1 {
		return Ops.Dasum(n, data, stride)
	}
	v := 0.0
	for i := 0; i < n; i++ {
		v += Ops.Dasum(m, data[i*stride:], 1)
	}
	return v
}

// Returns Y = X^T*Y
func DotD(X, Y MatDable) float64 {
	mX, nX := X.Size()
	mY, nY := Y.Size()
	if nX != 1 || nY != 1 || mX != mY {
		return math.NaN()
	}
	Xa := X.ArrayD()
	Ya := Y.ArrayD()
	return Ops.Ddot(mX, Xa, 1, Ya, 1)
}

func Max(X MatDable, dim int) (*MatD, []int) {
	m, n := X.Size()
	stride := X.Stride()
	data := X.ArrayD()

	var res *MatD
	var ixs []int
	if dim == 1 {
		res = Zeros(1, n)
		ixs = make([]int, n)
		for i := range ixs {
			ixs[i] = Ops.Idamax(n, data[i:], stride) - 1
			res.Set(data[i+ixs[i]*stride], i)
		}
	} else if dim == 0 {
		res = Zeros(m, 1)
		ixs = make([]int, m)
		for i := range ixs {
			ixs[i] = Ops.Idamax(m, data[i*stride:], 1) - 1
			res.Set(data[i*stride+ixs[i]], i)
		}
	} else {
		panic("dim must be 0 or 1")
	}
	return res, ixs
}
