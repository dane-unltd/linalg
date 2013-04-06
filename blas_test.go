package linalg

import (
	_ "github.com/dane-unltd/linalg/blasinit"
	"github.com/dane-unltd/linalg/matrix"
	"math"
	"testing"
)

func Benchmark_MulDelayed(b *testing.B) {
	A := matrix.RandN(1000, 1000)
	B := matrix.RandN(1000, 1000)
	C := matrix.RandN(1000, 1)
	for i := 0; i < b.N; i++ {
		matrix.MulD(A, B, C)
	}
}

func Benchmark_MulInstant(b *testing.B) {
	A := matrix.RandN(1000, 1000)
	B := matrix.RandN(1000, 1000)
	C := matrix.RandN(1000, 1)
	for i := 0; i < b.N; i++ {
		matrix.MulD(matrix.MulD(A, B), C)
	}
}

func Test_MulD(t *testing.T) {
	A := matrix.RandN(2, 600)
	B := matrix.RandN(600, 3000)
	C := matrix.RandN(3000, 100)
	D := matrix.RandN(100, 1000)
	E := matrix.RandN(1000, 10)
	R1 := matrix.MulD(matrix.MulD(matrix.MulD(matrix.MulD(A, B), C), D), E)
	R2 := matrix.MulD(A, B, C, D, E)

	if math.Abs(matrix.Nrm2D(R1)-matrix.Nrm2D(R2)) > 0.0001 {
		t.Error("MulD failed", matrix.Nrm2D(R1), matrix.Nrm2D(R2))
	}
}
