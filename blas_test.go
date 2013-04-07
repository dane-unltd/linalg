package linalg

import (
	"fmt"
	_ "github.com/dane-unltd/linalg/blasinit"
	. "github.com/dane-unltd/linalg/matrix"
	"math"
	"testing"
)

func Benchmark_MulDelayed(b *testing.B) {
	A := RandN(1000, 1000)
	B := RandN(1000, 1000)
	C := RandN(1000, 1)
	for i := 0; i < b.N; i++ {
		MulD(A, B, C)
	}
}

func Benchmark_MulInstant(b *testing.B) {
	A := RandN(1000, 1000)
	B := RandN(1000, 1000)
	C := RandN(1000, 1)
	for i := 0; i < b.N; i++ {
		MulD(MulD(A, B), C)
	}
}

func Test_MulD(t *testing.T) {
	A := RandN(2, 600)
	B := RandN(600, 3000)
	C := RandN(3000, 100)
	D := RandN(100, 1000)
	E := RandN(1000, 10)
	R1 := MulD(MulD(MulD(MulD(A, B), C), D), E)
	R2 := MulD(A, B, C, D, E)

	if math.Abs(Nrm2D(R1)-Nrm2D(R2)) > 0.0001 {
		t.Error("MulD failed", Nrm2D(R1), Nrm2D(R2))
	}

	A = FromArrayD([]float64{1, -1, 0, 1}, true, 2, 2)

	fmt.Println(A.SolveTriU(VecD{1, 1}))
	fmt.Println(A.SolveTriL(VecD{1, 1}))
	A.Tr()
	fmt.Println(A.SolveTriU(VecD{1, 1}))
	fmt.Println(A.SolveTriL(VecD{1, 1}))
	fmt.Println(A.IsTr())
}
