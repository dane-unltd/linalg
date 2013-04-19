package linalg

import (
	"github.com/dane-unltd/linalg/lapackimpl"
	. "github.com/dane-unltd/linalg/matrix"
	"testing"
)

func TestSvd(t *testing.T) {
	A := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)
	S := NewDiagFloat64(2)
	U := NewDenseFloat64(2, 2)
	Vt := NewDenseFloat64(2, 2)

	lapack.Dgsvd(A, S, U, Vt)

	A2 := NewDenseFloat64(2, 2)
	A3 := NewDenseFloat64(2, 2)
	A2.Mul(U, S)
	A3.Mul(A2, Vt)

	A2.Sub(A3, A)

	if A2.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", A3)
	}
}
