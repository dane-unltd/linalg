package matrix

import "math"

//level1
func (x VecD) Nrm2() float64 {
	return Ops.Dnrm2(len(x), x, 1)
}

func (x VecD) Asum() float64 {
	return Ops.Dasum(len(x), x, 1)
}

func (x VecD) AddaY(y VecD, arr ...float64) VecD {
	a := 1.0
	if len(arr) > 0 {
		a = arr[0]
	}
	Ops.Daxpy(len(x), a, y, 1, x, 1)
	return x
}

func (x VecD) Scal(a float64) VecD {
	Ops.Dscal(len(x), a, x, 1)
	return x
}

func (res VecD) Mul(a float64, x VecD) VecD {
	for i := range res {
		res[i] = a * x[i]
	}
	return res
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

func (x VecD) Idamax() int {
	return Ops.Idamax(len(x), x, 1)
}

func (x VecD) Idmin() int {
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

func (x VecD) Idmax() int {
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

func (x VecD) Nrm2Sq() float64 {
	res := 0.0
	for _, v := range x {
		res += v * v
	}
	return res
}

func (res VecD) Normalize(x VecD) VecD {
	return res.Mul(1/x.Nrm2(), x)
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

func (res VecD) Cross(x, y VecD) VecD {
	res[0] = x[1]*y[2] - x[2]*y[1]
	res[1] = x[2]*y[0] - x[0]*y[2]
	res[2] = x[0]*y[1] - x[1]*y[0]
	return res
}
