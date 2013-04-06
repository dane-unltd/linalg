package linalg

import (
	_ "github.com/dane-unltd/linalg/blasinit"
	_ "github.com/dane-unltd/linalg/lapackinit"
	"github.com/dane-unltd/linalg/matrix"
	"math"
	"testing"
)

func Test_SvdD(t *testing.T) {
	A1 := matrix.RandN(30, 40)

	S, U, Vt := matrix.SvdD(A1)
	A2 := matrix.MulD(S, U, Vt)
	if math.Abs(matrix.Nrm2D(A1)-matrix.Nrm2D(A2)) > 0.0001 {
		t.Error("SvdD failed", matrix.Nrm2D(A1), matrix.Nrm2D(A2))
	}
}
