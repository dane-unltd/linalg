package matrix

type VecFloat64 []float64

func NewVecFloat64(n int) VecFloat64 {
	return make(VecFloat64, n)
}

func (v VecFloat64) Size() (int, int) {
	return len(v), 1
}

func (v VecFloat64) At(i, j int) float64 {
	if j != 0 {
		panic("j has to be 0")
	}
	return v[i]
}

func (v VecFloat64) Copy() interface{} {
	vNew := make(VecFloat64, len(v))
	copy(vNew, v)
	return vNew
}

func (v VecFloat64) Equals(x interface{}) bool {

	M, ok := x.(Matrix)
	if !ok {
		return ok
	}

	m, n := M.Size()
	if n > 1 || m != len(v) {
		return false
	}
	for i := range v {
		if v[i] != M.At(i, 0) {
			return false
		}
	}
	return true
}
