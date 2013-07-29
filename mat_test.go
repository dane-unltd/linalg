package linalg

import (
	"github.com/dane-unltd/linalg/clapack"
	"github.com/dane-unltd/linalg/mat"
	"github.com/kortschak/cblas"
	"testing"
)

var n = 3

type cblasops struct {
	cblas.Blas
	clapack.Lapack
}

func TestMul(t *testing.T) {
	mat.Register(cblasops{})
	A := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	B := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)

	res := &mat.Dense{}

	resCorrect := mat.NewFromArray([]float64{7, 10, 15, 22}, false, 2, 2)

	res.Mul(A, B)

	temp := mat.Dense{}
	temp.Sub(res, resCorrect)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	resCorrect = mat.NewFromArray([]float64{3, 6, 4, 8}, false, 2, 2)

	Asub := A.View(0, 0, 2, 1)
	Bsub := B.View(0, 1, 2, 1).TrView()

	res.Mul(Asub, Bsub)

	temp.Sub(res, resCorrect)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	Bsub = B.TrView().View(1, 0, 1, 2)

	res.Mul(Asub, Bsub)

	temp.Sub(res, resCorrect)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	//View as receiver:
	res = mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	dst := res.View(0, 0, 2, 1)

	resCorrect = mat.NewFromArray([]float64{7, 10, 3, 4}, false, 2, 2)

	dst.Mul(B, Asub)

	temp.Sub(res, resCorrect)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	res = mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	dst = res.TrView().View(0, 0, 2, 1)

	resCorrect = mat.NewFromArray([]float64{7, 2, 10, 4}, false, 2, 2)

	dst.Mul(B, Asub)

	temp.Sub(res, resCorrect)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}
}

func TestAdd(t *testing.T) {
	mat.Register(cblasops{})
	A := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	B := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)

	res := &mat.Dense{}

	resCorrect := mat.NewFromArray([]float64{2, 4, 6, 8}, false, 2, 2)

	res.Add(A, B)

	temp := mat.Dense{}
	temp.Sub(res, resCorrect)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	resCorrect = mat.NewFromArray([]float64{2, 5}, false, 2, 1)

	Asub := A.View(0, 0, 2, 1)
	Bsub := B.View(0, 0, 1, 2).TrView()

	res.Add(Asub, Bsub)

	temp.Sub(res, resCorrect)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	Bsub = B.TrView().View(0, 0, 2, 1)

	res.Add(Asub, Bsub)

	temp.Sub(res, resCorrect)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

}
