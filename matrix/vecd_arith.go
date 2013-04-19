package matrix

import "github.com/dane-unltd/linalg/blas"

func (res VecD) Normalize(v VecD) VecD {
	blas.Dcopy(len(res), v, 1, res, 1)
	return res.Scal(1 / v.Nrm2())
}

func (v VecD) Nrm2Sq() float64 {
	return blas.Ddot(len(v), v, 1, v, 1)
}

func (v VecD) Nrm2() float64 {
	return blas.Dnrm2(len(v), v, 1)
}

func (v VecD) Asum() float64 {
	return blas.Dasum(len(v), v, 1)
}

func (v VecD) Imax() int {
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

func (res VecD) Max(v VecD, a float64) VecD {
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

func (res VecD) MulH(a, b VecD) {
	if len(res) != len(a) || len(res) != len(b) {
		panic("dimension missmatch")
	}
	for i := range res {
		res[i] = a[i] * b[i]
	}
}

func (res VecD) Scal(a float64) VecD {
	blas.Dscal(len(res), a, res, 1)
	return res
}

func (res VecD) Axpy(a float64, x VecD) VecD {
	blas.Daxpy(len(res), a, x, 1, res, 1)
	return res
}

func Ddot(a, b VecD) float64 {
	if len(a) != len(b) {
		panic("dimension missmatch")
	}
	return blas.Ddot(len(a), a, 1, b, 1)
}

func (res VecD) Add(a, b VecD) VecD {
	if len(res) != len(a) || len(res) != len(b) {
		panic("dimension missmatch")
	}

	for i := range res {
		res[i] = a[i] + b[i]
	}
	return res
}

func (res VecD) Sub(a, b VecD) VecD {
	if len(res) != len(a) || len(res) != len(b) {
		panic("dimension missmatch")
	}

	for i := range res {
		res[i] = a[i] - b[i]
	}
	return res
}

func (res VecD) Cross(a, b VecD) VecD {
	if len(res) != 3 || len(a) != 3 || len(b) != 3 {
		panic("All vectors must have length 3")
	}
	res[0] = a[1]*b[2] - a[2]*b[1]
	res[1] = a[2]*b[0] - a[0]*b[2]
	res[2] = a[0]*b[1] - a[1]*b[0]
	return res
}

func (res VecD) Neg(v VecD) VecD {
	if len(res) != len(v) {
		panic("dimension missmatch")
	}
	for i := range res {
		res[i] = -v[i]
	}
	return res
}

func (res VecD) Mul(A, B Matrix) {
	resMat := FromArrayD(res, true, len(res), 1)
	resMat.Mul(A, B)
}
