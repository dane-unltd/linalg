package linalg

import (
	"github.com/dane-unltd/linalg/mat"
	"testing"
)

func TestSvd(t *testing.T) {
	mat.Register(goblasops{})
	A := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	S := mat.NewVec(2)
	U := mat.NewDense(2, 2)
	Vt := mat.NewDense(2, 2)

	A.Svd(S, U, Vt, false)

	A2 := mat.NewDense(2, 2)
	A3 := mat.NewDense(2, 2)
	A2.ScalCols(U, S)
	A3.Mul(A2, Vt)

	A2.Sub(A3, A)

	if A2.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", A3)
	}
}
func TestChol(t *testing.T) {
	mat.Register(goblasops{})
	A := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	At := A.TrView()
	AtA := mat.NewDense(2)
	AtA.Mul(At, A)

	R := mat.NewDense(2)
	Rt := R.TrView()

	AtA.Chol(R)

	A2 := mat.NewDense(2)
	A2.Mul(Rt, R)

	A3 := mat.NewDense(2)
	A3.Sub(A2, AtA)

	if A3.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", A3)
	}
}
