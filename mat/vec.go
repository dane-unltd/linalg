package mat

import "math/rand"

type Vec []float64

func NewVec(n int) Vec {
	return make(Vec, n)
}

func RandVec(n int) Vec {
	v := NewVec(n)
	for i := range v {
		v[i] = rand.Float64()
	}
	return v
}

func (v Vec) Dims() (int, int) {
	return len(v), 1
}

func (v Vec) At(i, j int) float64 {
	if j != 0 {
		panic("j has to be 0")
	}
	return v[i]
}

func (dst Vec) Copy(src Vec) Vec {
	if len(dst) != len(src) {
		panic("dimension missmatch")
	}
	ops.Dcopy(len(dst), src, 1, dst, 1)
	return dst
}

func (v Vec) Equals(a Vec) bool {
	if len(a) != len(v) {
		return false
	}

	for i := range v {
		if v[i] != a[i] {
			return false
		}
	}
	return true
}

func (v Vec) Col(ix int, c Vec) Vec {
	if c == nil {
		return v
	}

	c.Copy(v)
	return c
}
