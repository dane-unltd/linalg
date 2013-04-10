package linalg

import (
	"github.com/dane-unltd/linalg/lapack"
	. "github.com/dane-unltd/linalg/matrix"
	"testing"
)

func TestSvd(t *testing.T) {
	A := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)
	S := NewDiagD(2)
	U := NewDenseD(2, 2)
	Vt := NewDenseD(2, 2)

	lapack.DenseDSvd(A, S, U, Vt)

	A2 := NewDenseD(2, 2)
	A3 := NewDenseD(2, 2)
	A2.Mul(U, S)
	A3.Mul(A2, Vt)

	A2.Sub(A3, A)

	if A2.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", A3)
	}
}
