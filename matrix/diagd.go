package matrix

type DiagD []float64

func NewDiagD(m int) DiagD {
	return make(DiagD, m)
}

func (D DiagD) At(i, j int) float64 {
	if i == j {
		return D[i]
	}
	return 0
}

func (D DiagD) Set(i, j int, v float64) {
	if i != j {
		panic("i!=j")
	}
	D[i] = v
}

func (D DiagD) Size() (int, int) {
	return len(D), len(D)
}
