package matrix

import "math"

var VecDNormalize = VecDNormalizeGo

func VecDNormalizeGo(res, v VecD) VecD {
	return res.Scale(v, 1/v.Nrm2())
}
func (res VecD) Normalize(v VecD) VecD {
	return VecDNormalize(res, v)
}

var VecDNrm2Sq = VecDNrm2SqGo

func VecDNrm2SqGo(v VecD) float64 {
	res := 0.0
	for _, val := range v {
		res += val * val
	}
	return res
}
func (v VecD) Nrm2Sq() float64 {
	return VecDNrm2Sq(v)
}

var VecDNrm2 = VecDNrm2Go

func VecDNrm2Go(v VecD) float64 {
	return math.Sqrt(v.Nrm2Sq())
}
func (v VecD) Nrm2() float64 {
	return VecDNrm2(v)
}

var VecDAsum = VecDAsumGo

func VecDAsumGo(v VecD) float64 {
	res := 0.0
	for _, val := range v {
		res += math.Abs(val)
	}
	return res
}
func (v VecD) Asum() float64 {
	return VecDAsum(v)
}

func (v VecD) Idmax() int {
	max := 0.0
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

var VecDScale = VecDScaleGo

func VecDScaleGo(res, v VecD, a float64) VecD {
	if len(res) != len(v) {
		panic("dimension missmatch")
	}
	for i := range res {
		res[i] = a * v[i]
	}
	return res
}
func (res VecD) Scale(v VecD, a float64) VecD {
	return VecDScale(res, v, a)
}

var Dot = DotGo

func DotGo(a, b VecD) float64 {
	if len(a) != len(b) {
		panic("dimension missmatch")
	}
	res := 0.0
	for i := range a {
		res = a[i] * b[i]
	}
	return res
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
