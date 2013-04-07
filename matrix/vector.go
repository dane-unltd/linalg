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

func (v VecD) IsTr() bool {
	return false
}

func ZeroVec(n int) VecD {
	return make(VecD, n)
}

func OnesVec(n int) VecD {
	v := make(VecD, n)
	for i := range v {
		v[i] = 1
	}
	return v
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

func (x VecD) MinIx() int {
	min := x[0]
	minIx := 0
	for i, v := range x {
		if v < min {
			min = v
			minIx = i
		}
	}
	return minIx
}

func (x VecD) MaxIx() int {
	max := x[0]
	maxIx := 0
	for i, v := range x {
		if v > max {
			max = v
			maxIx = i
		}
	}
	return maxIx
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

func (x VecD) Norm1() float64 {
	sum := 0.0
	for _, v := range x {
		sum += v
	}
	return sum
}

func (x VecD) Norm2Sq() float64 {
	sum := 0.0
	for _, v := range x {
		sum += v * v
	}
	return sum
}

func (x VecD) Norm2() float64 {
	return math.Sqrt(x.Norm2Sq())
}

func (res VecD) Normalize(x VecD) VecD {
	return res.Mul(1/x.Norm2(), x)
}

func (x VecD) Dot(y VecD) float64 {
	if len(x) != len(y) {
		panic("dimension missmatch")
	}
	sum := 0.0
	for i := range x {
		sum += x[i] * y[i]
	}
	return sum
}
