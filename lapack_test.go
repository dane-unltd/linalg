package linalg

import (
	"github.com/dane-unltd/linalg/clapack"
	. "github.com/dane-unltd/linalg/matrix"
	"testing"
)

func TestSvd(t *testing.T) {
	A := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)
	S := NewDiag(2)
	U := NewDense(2, 2)
	Vt := NewDense(2, 2)

	lapack.Dgsvd(A, S, U, Vt)

	A2 := NewDense(2, 2)
	A3 := NewDense(2, 2)
	A2.Mul(U, S)
	A3.Mul(A2, Vt)

	A2.Sub(A3, A)

	if A2.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", A3)
	}
}
