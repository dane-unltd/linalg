package linalg

import (
	"github.com/dane-unltd/linalg/matrix"
	"testing"
)

func TestSvd(t *testing.T) {
	matrix.Register(goblasops{})
	A := matrix.FromArray([]float64{1, 2, 3, 4}, true, 2, 2)
	S := matrix.NewDiag(2)
	U := matrix.NewDense(2, 2)
	Vt := matrix.NewDense(2, 2)

	A.Svd(S, U, Vt)

	A2 := matrix.NewDense(2, 2)
	A3 := matrix.NewDense(2, 2)
	A2.Mul(U, S)
	A3.Mul(A2, Vt)

	A2.Sub(A3, A)

	if A2.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", A3)
	}
}
