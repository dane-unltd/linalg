package matrix

type Diag []float64

func NewDiag(m int) Diag {
	return make(Diag, m)
}

func (D Diag) At(i, j int) float64 {
	if i > len(D) || j > len(D) {
		panic("index out of bounds")
	}
	if i == j {
		return D[i]
	}
	return 0
}

func (D Diag) Set(i, j int, v float64) {
	if i != j {
		panic("i!=j")
	}
	D[i] = v
}

func (D Diag) Size() (int, int) {
	return len(D), len(D)
}

func (dst Diag) Copy(src Diag) {
	if len(dst) != len(src) {
		panic("dimension missmatch")
	}
	copy(dst, src)
}
