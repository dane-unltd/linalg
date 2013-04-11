package linalg

import (
	"fmt"
	_ "github.com/dane-unltd/linalg/blas"
	. "github.com/dane-unltd/linalg/matrix"
	"github.com/harrydb/go/matrix"
	"github.com/ziutek/blas"
	_ "math"
	_ "math/rand"
	"testing"
)

func Benchmark_MatrixMul(b *testing.B) {
	b.StopTimer()
	n := 2000
	A := RandN(n, n)
	B := RandN(n, n)
	res := RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_MatrixMulBlasGo(b *testing.B) {
	b.StopTimer()
	n := 2000
	A := RandN(n, n)
	B := RandN(n, n)
	Am := matrix.New(n, n, A.VecView())
	Bm := matrix.New(n, n, B.VecView())
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		matrix.MulDouglas(Am, Bm)
	}
}

func Benchmark_Nrm2Go(b *testing.B) {
	A := RandN(100, 100)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}
func Benchmark_Nrm2GoFast(b *testing.B) {
	A := RandN(100, 100)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		blas.Dnrm2(len(v), v, 1)
	}
}

func TestMatrixBlas(t *testing.T) {
	A := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)

	D := DiagD{3, 4}

	res := NewDenseD(2, 2)

	res.Mul(A, D)

	fmt.Println(res)
}
