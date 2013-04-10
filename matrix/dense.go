package matrix

type dense struct {
	rows, cols int
	stride     int
	trans      bool
}

func (D *dense) Size() (int, int) {
	if D.trans {
		return D.cols, D.rows
	}
	return D.rows, D.cols
}

func (D *dense) Stride() int {
	return D.stride
}

func (D *dense) IsTr() bool {
	return D.trans
}
