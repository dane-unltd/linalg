package matrix

import "math"

type VecD []float64

func (v VecD) Size() (int, int) {
	return len(v), 1
}
func (v VecD) Stride() int {
	return len(v)
}
func (v VecD) ArrayD() []float64 {
	return []float64(v)
}

func ZeroVec(n int) VecD {
	return make(VecD, n)
}

func (v VecD) Copy() VecD {
	vNew := make(VecD, len(v))
	copy(vNew, v)
	return vNew
}

func (v VecD) Equals(M MatDable) bool {
	m, n := M.Size()
	if n > 1 || m != len(v) {
		return false
	}
	v2 := M.ArrayD()
	for i := range v {
		if v[i] != v2[i] {
			return false
		}
	}
	return true
}

func (res VecD) Add(x, y VecD) VecD {
	for i := range res {
		res[i] = x[i] + y[i]
	}
	return res
}

func (res VecD) Sub(x, y VecD) VecD {
	for i := range res {
		res[i] = x[i] - y[i]
	}
	return res
}

func (res VecD) Neg(y VecD) VecD {
	for i := range res {
		res[i] = -y[i]
	}
	return res
}

func (res VecD) Mul(a float64, x VecD) VecD {
	for i := range res {
		res[i] = a * x[i]
	}
	return res
}

func (res VecD) MulH(x, y VecD) VecD {
	for i := range res {
		res[i] = x[i] * y[i]
	}
	return res
}

func (res VecD) Max(x VecD, m float64) VecD {
	for i := range res {
		res[i] = math.Max(x[i], m)
	}
	return res
}

func (res VecD) Abs(x VecD) VecD {
	for i := range res {
		res[i] = math.Abs(x[i])
	}
	return res
}

func (res VecD) Sign(x VecD) VecD {
	for i := range res {
		if x[i] > 0 {
			res[i] = 1
		} else if x[i] < 0 {
			res[i] = -1
		} else {
			res[i] = 0
		}
	}
	return res
}
