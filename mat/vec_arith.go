package mat

import "github.com/gonum/blas"

func (res Vec) Normalize(v Vec) Vec {
	ops.Dcopy(len(res), v, 1, res, 1)
	return res.Scal(1 / v.Nrm2())
}

func (v Vec) Nrm2Sq() float64 {
	return ops.Ddot(len(v), v, 1, v, 1)
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

func (res Vec) Mul(a Vec) Vec {
	if len(res) != len(a) {
		panic("dimension missmatch")
	}
	for i, v := range a {
		res[i] *= v
	}
	return res
}

func (res Vec) Div(a Vec) Vec {
	if len(res) != len(a) {
		panic("dimension missmatch")
	}
	for i, v := range a {
		res[i] /= v
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

func (res Vec) Add(a Vec) Vec {
	if len(res) != len(a) {
		panic("dimension missmatch")
	}

	for i, v := range a {
		res[i] += v
	}
	return res
}

func (res Vec) Sub(a Vec) Vec {
	if len(res) != len(a) {
		panic("dimension missmatch")
	}

	for i, v := range a {
		res[i] -= v
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

func (res Vec) Neg() Vec {
	for i, v := range res {
		res[i] = -v
	}
	return res
}

func (res Vec) AddMul(A Matrix, x Vec, a float64) {
	m, n := A.Dims()
	if len(res) != m || len(x) != n {
		panic("dimension missmatch")
	}

	switch A := A.(type) {
	case *Dense:
		ops.Dgemv(blas.ColMajor, blas.NoTrans, m, n, a, A.data, A.rows, x, 1,
			1, res, 1)
		return
	case *denseView:
		if A.trans == blas.Trans {
			m, n = n, m
		}
		ops.Dgemv(blas.ColMajor, A.trans, m, n, a, A.data, A.stride, x, 1,
			1, res, 1)
		return
	}
	panic("general mv not implemented")
}

func (dst Vec) MMul(A, v Matrix) Vec {
	m, n := A.Dims()
	mv, nv := v.Dims()
	if mv != n || len(dst) != m {
		panic("dimension mismatch")
	}
	if nv != 1 {
		panic("right hand side has to be a vector")
	}

	if v, ok := v.(Coler); ok {
		vc := v.Col(1, nil)
		switch A := A.(type) {
		case *Dense:
			ops.Dgemv((blas.ColMajor), (blas.NoTrans), m, n, 1,
				A.data, m, vc, 1, 0, dst, 1)
			return dst
		case *denseView:
			if A.IsTr() {
				m, n = n, m
			}
			ops.Dgemv((blas.ColMajor), A.trans, m, n, 1,
				A.data, m, vc, 1, 0, dst, 1)
			return dst
		}
		for row := 0; row < m; row++ {
			sum := 0.0
			for col := 0; col < n; col++ {
				sum += A.At(row, col) * vc[col]
			}
			dst[row] = sum
		}
		return dst
	}

	for row := 0; row < m; row++ {
		sum := 0.0
		for col := 0; col < n; col++ {
			sum += A.At(row, col) * v.At(col, 1)
		}
		dst[row] = sum
	}
	return dst
}
