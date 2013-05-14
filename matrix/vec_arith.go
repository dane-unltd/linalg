package matrix

import "github.com/kortschak/blas"

func (res Vec) Normalize(v Vec) Vec {
	ops.Dcopy(len(res), v, 1, res, 1)
	return res.Scal(1 / v.Nrm2())
}

func (v Vec) Nrm2Sq() float64 {
	return ops.Ddot(len(v), v, 1, v, 1)
}

func (res Vec) CopyFrom(v Vec) {
	ops.Dcopy(len(res), v, 1, res, 1)
}

func (v Vec) Nrm2() float64 {
	return ops.Dnrm2(len(v), v, 1)
}

func (v Vec) Asum() float64 {
	return ops.Dasum(len(v), v, 1)
}

func (v Vec) Imax() int {
	max := v[0]
	ixMax := 0
	for i, val := range v {
		if val > max {
			max = val
			ixMax = i
		}
	}
	return ixMax
}

func (res Vec) Max(v Vec, a float64) Vec {
	if len(res) != len(v) {
		panic("dimension missmatch")
	}
	for i, val := range v {
		if val > a {
			res[i] = val
		} else {
			res[i] = a
		}
	}
	return res
}

func (res Vec) MulH(a, b Vec) Vec {
	if len(res) != len(a) || len(res) != len(b) {
		panic("dimension missmatch")
	}
	for i := range res {
		res[i] = a[i] * b[i]
	}
	return res
}

func (res Vec) DivH(a, b Vec) Vec {
	if len(res) != len(a) || len(res) != len(b) {
		panic("dimension missmatch")
	}
	for i := range res {
		res[i] = a[i] / b[i]
	}
	return res
}

func (v Vec) AddSc(a float64) Vec {
	for i := range v {
		v[i] += a
	}
	return v
}

func (res Vec) Scal(a float64) Vec {
	ops.Dscal(len(res), a, res, 1)
	return res
}

func (res Vec) Axpy(a float64, x Vec) Vec {
	ops.Daxpy(len(res), a, x, 1, res, 1)
	return res
}

func Dot(a, b Vec) float64 {
	if len(a) != len(b) {
		panic("dimension missmatch")
	}
	return ops.Ddot(len(a), a, 1, b, 1)
}

func (res Vec) Add(a, b Vec) Vec {
	if len(res) != len(a) || len(res) != len(b) {
		panic("dimension missmatch")
	}

	for i := range res {
		res[i] = a[i] + b[i]
	}
	return res
}

func (res Vec) Sub(a, b Vec) Vec {
	if len(res) != len(a) || len(res) != len(b) {
		panic("dimension missmatch")
	}

	for i := range res {
		res[i] = a[i] - b[i]
	}
	return res
}

func (res Vec) Cross(a, b Vec) Vec {
	if len(res) != 3 || len(a) != 3 || len(b) != 3 {
		panic("All vectors must have length 3")
	}
	res[0] = a[1]*b[2] - a[2]*b[1]
	res[1] = a[2]*b[0] - a[0]*b[2]
	res[2] = a[0]*b[1] - a[1]*b[0]
	return res
}

func (res Vec) Neg(v Vec) Vec {
	if len(res) != len(v) {
		panic("dimension missmatch")
	}
	for i := range res {
		res[i] = -v[i]
	}
	return res
}

func (res Vec) Mul(A, B Matrix) {
	resMat := NewFromArray(res, false, len(res), 1)
	resMat.Mul(A, B)
}

func (res Vec) AddMul(A *Dense, x Vec, a float64) {
	m, n := A.Size()
	ops.Dgemv(blas.ColMajor, A.trans, m, n, a, A.data, A.stride, x, 1,
		1, res, 1)
}
