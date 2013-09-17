package linalg

import (
	"github.com/dane-unltd/linalg/mat"
	"testing"
)

func TestMMul(t *testing.T) {
	mat.Register(cblasops{})
	A := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	B := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)

	res := &mat.Dense{}

	resCorrect := mat.NewFromArray([]float64{7, 10, 15, 22}, false, 2, 2)

	res.MMul(A, B)

	temp := mat.Dense{}
	temp.Copy(res).Sub(resCorrect)
	if temp.Vec(nil).Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	resCorrect = mat.NewFromArray([]float64{3, 6, 4, 8}, false, 2, 2)

	Asub := A.View(0, 0, 2, 1)
	Bsub := B.View(0, 1, 2, 1).TrView()

	res.MMul(Asub, Bsub)

	temp.Copy(res).Sub(resCorrect)
	if temp.Vec(nil).Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	Bsub = B.TrView().View(1, 0, 1, 2)

	res.MMul(Asub, Bsub)

	temp.Copy(res).Sub(resCorrect)
	if temp.Vec(nil).Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	//View as receiver:
	res = mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	dst := res.View(0, 0, 2, 1)

	resCorrect = mat.NewFromArray([]float64{7, 10, 3, 4}, false, 2, 2)

	dst.MMul(B, Asub)

	temp.Copy(res).Sub(resCorrect)
	if temp.Vec(nil).Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

}

func TestAdd(t *testing.T) {
	mat.Register(cblasops{})
	A := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	B := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)

	res := &mat.Dense{}

	resCorrect := mat.NewFromArray([]float64{2, 4, 6, 8}, false, 2, 2)

	res.Copy(A).Add(B)

	temp := mat.Dense{}
	temp.Copy(res).Sub(resCorrect)
	if temp.Vec(nil).Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	resCorrect = mat.NewFromArray([]float64{2, 5}, false, 2, 1)

	Asub := A.View(0, 0, 2, 1)
	Bsub := B.View(0, 0, 1, 2).TrView()

	res.Copy(Asub).Add(Bsub)

	temp.Copy(res).Sub(resCorrect)
	if temp.Vec(nil).Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

	Bsub = B.TrView().View(0, 0, 2, 1)

	res.Copy(Asub).Add(Bsub)

	temp.Copy(res).Sub(resCorrect)
	if temp.Vec(nil).Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, resCorrect)
	}

}
