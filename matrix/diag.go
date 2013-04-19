package matrix

type DiagFloat64 []float64

func NewDiagFloat64(m int) DiagFloat64 {
	return make(DiagFloat64, m)
}

func (D DiagFloat64) At(i, j int) float64 {
	if i == j {
		return D[i]
	}
	return 0
}

func (D DiagFloat64) Set(i, j int, v float64) {
	if i != j {
		panic("i!=j")
	}
	D[i] = v
}

func (D DiagFloat64) Size() (int, int) {
	return len(D), len(D)
}

func (D DiagFloat64) Copy() Matrix {
	dNew := make(DiagFloat64, len(D))
	copy(dNew, D)
	return dNew
}
