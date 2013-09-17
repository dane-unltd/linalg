package linalg

import (
	"github.com/dane-unltd/linalg/mat"
	"testing"
)

func TestSvd(t *testing.T) {
	mat.Register(goblasops{})
	A := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	S := mat.NewVec(2)
	U := mat.New(2, 2)
	Vt := mat.New(2, 2)

	A.Svd(S, U, Vt, false)
	t.Log(U, S)

	A2 := mat.New(2, 2)
	A2.Copy(U).ScalCols(S)
	A3 := mat.New(2, 2)
	A3.MMul(A2, Vt)

	A2.Copy(A3).Sub(A)

	if A2.Vec(nil).Nrm2Sq() > 0.01 {
		t.Error("wrong result", A3)
	}
}

func TestChol(t *testing.T) {
	mat.Register(goblasops{})
	A := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	At := A.TrView()
	AtA := mat.New(2)
	AtA.MMul(At, A)

	R := mat.New(2)
	Rt := R.TrView()

	AtA.Chol(R)

	A2 := mat.New(2)
	A2.MMul(Rt, R)

	A3 := mat.New(2)
	A3.Copy(A2).Sub(AtA)

	if A3.Vec(nil).Nrm2Sq() > 0.01 {
		t.Error("wrong result", A3)
	}
}
