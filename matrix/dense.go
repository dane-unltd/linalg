package matrix

import blas "github.com/dane-unltd/linalg/blasops"

type dense struct {
	rows, cols int
	stride     int
	trans      blas.Transpose
}

func (D *dense) Size() (int, int) {
	if D.IsTr() {
		return D.cols, D.rows
	}
	return D.rows, D.cols
}

func (D *dense) Stride() int {
	return D.stride
}

func (D *dense) IsTr() bool {
	return D.trans == blas.Trans
}
