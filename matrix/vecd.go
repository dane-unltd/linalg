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

func (v VecD) Copy() Matrix {
	vNew := make(VecD, len(v))
	copy(vNew, v)
	return vNew
}

func (v VecD) Equals(M Matrix) bool {
	m, n := M.Size()
	if n > 1 || m != len(v) {
		return false
	}
	for i := range v {
		if v[i] != M.At(i, 1) {
			return false
		}
	}
	return true
}
