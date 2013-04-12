package linalg

import (
	cblas "code.google.com/p/biogo.blas"
	"fmt"
	"github.com/dane-unltd/linalg/blas"
	. "github.com/dane-unltd/linalg/matrix"
	goblas "github.com/ziutek/blas"
	"testing"
)

var n = 50

func Benchmark_MatrixMulGo(b *testing.B) {
	b.StopTimer()
	A := RandN(n, n)
	B := RandN(n, n)
	res := RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_MatrixMulBlas(b *testing.B) {
	b.StopTimer()
	A := RandN(n, n)
	B := RandN(n, n)
	res := RandN(n, n)
	blas.Dgemm = cblas.Dgemm
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}
func Benchmark_MatrixMulBlasGo(b *testing.B) {
	b.StopTimer()
	A := RandN(n, n)
	B := RandN(n, n)
	res := RandN(n, n)
	blas.Dgemm = goblas.Dgemm
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_Nrm2Go(b *testing.B) {
	A := RandN(100, 100)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}
func Benchmark_Nrm2Blas(b *testing.B) {
	A := RandN(100, 100)
	v := A.VecView()
	blas.Dnrm2 = cblas.Dnrm2

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}
func Benchmark_Nrm2GoFast(b *testing.B) {
	A := RandN(100, 100)
	v := A.VecView()
	blas.Dnrm2 = goblas.Dnrm2

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func TestMatrixBlas(t *testing.T) {
	A := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)

	D := DiagD{3, 4}

	res := NewDenseD(2, 2)

	res.Mul(A, D)

	fmt.Println(res)
}
