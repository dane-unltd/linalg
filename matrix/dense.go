package matrix

type dense struct {
	rows, cols int
	stride     int
}

func (D *dense) Size() (int, int) {
	return D.rows, D.cols
}

func (D *dense) Stride() int {
	return D.stride
}
