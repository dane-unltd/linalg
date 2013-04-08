package linalg

import (
	"fmt"
	_ "github.com/dane-unltd/linalg/blasinit"
	_ "github.com/dane-unltd/linalg/lapackinit"
	"github.com/dane-unltd/linalg/matrix"
	"math"
	"testing"
)

func Test_SvdD(t *testing.T) {
	A1 := matrix.FromArrayD([]float64{1, 2, 3, 4, 5, 6}, true, 3, 2)

	S, U, Vt := matrix.SvdD(A1.Copy(), false)
	fmt.Println(U)
	fmt.Println(S)
	fmt.Println(Vt)
	A2 := matrix.MulD(U, S, Vt)
	fmt.Println(A1)
	fmt.Println(A2)
	if math.Abs(matrix.Nrm2D(A1)-matrix.Nrm2D(A2)) > 0.0001 {
		t.Error("SvdD failed", matrix.Nrm2D(A1), matrix.Nrm2D(A2))
	}
}
