package matrix

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
