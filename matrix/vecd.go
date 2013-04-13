package matrix

type VecD []float64

func NewVecD(n int) VecD {
	return make(VecD, n)
}

func (v VecD) Size() (int, int) {
	return len(v), 1
}

func (v VecD) At(i, j int) float64 {
	if j != 0 {
		panic("j has to be 0")
	}
	return v[i]
}

func (v VecD) Copy() interface{} {
	vNew := make(VecD, len(v))
	copy(vNew, v)
	return vNew
}

func (v VecD) Equals(x interface{}) bool {

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
