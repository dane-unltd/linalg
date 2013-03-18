package blas

import (
	"github.com/dane-unltd/linalg/dense"
	"math"
	"testing"
)

func Benchmark_MulDelayed(b *testing.B) {
	A := dense.RandN(1000, 1000)
	B := dense.RandN(1000, 1000)
	C := dense.RandN(1000, 1)
	for i := 0; i < b.N; i++ {
		MulD(A, B, C)
	}
}

func Benchmark_MulInstant(b *testing.B) {
	A := dense.RandN(1000, 1000)
	B := dense.RandN(1000, 1000)
	C := dense.RandN(1000, 1)
	for i := 0; i < b.N; i++ {
		MulD(MulD(A, B), C)
	}
}

func Test_MulD(t *testing.T) {
	A := dense.RandN(2, 600)
	B := dense.RandN(600, 3000)
	C := dense.RandN(3000, 100)
	D := dense.RandN(100, 1000)
	E := dense.RandN(1000, 10)
	R1 := MulD(MulD(MulD(MulD(A, B), C), D), E)
	R2 := MulD(A, B, C, D, E)

	if math.Abs(Nrm2D(R1)-Nrm2D(R2)) > 0.0001 {
		t.Error("MulD failed", Nrm2D(R1), Nrm2D(R2))
	}
}
