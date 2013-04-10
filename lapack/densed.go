package lapack

import "github.com/dane-unltd/linalg/matrix"

type DenseD struct {
	*matrix.DenseD
}

func NewDenseD(m, n int) DenseD {
	return DenseD{matrix.NewDenseD(m, n)}
}
