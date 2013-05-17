package matrix

type Vec []float64

func NewVec(n int) Vec {
	return make(Vec, n)
}

func (v Vec) Size() (int, int) {
	return len(v), 1
}

func (v Vec) At(i, j int) float64 {
	if j != 0 {
		panic("j has to be 0")
	}
	return v[i]
}

func (dst Vec) Copy(src Vec) {
	if len(dst) != len(src) {
		panic("dimension missmatch")
	}
	ops.Dcopy(len(dst), src, 1, dst, 1)
}

func (v Vec) Equals(x interface{}) bool {

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
