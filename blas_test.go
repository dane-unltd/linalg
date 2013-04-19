package linalg

import (
	cblas "code.google.com/p/biogo.blas"
	"fmt"
	"github.com/dane-unltd/linalg/blas"
	_ "github.com/dane-unltd/linalg/goblas"
	. "github.com/dane-unltd/linalg/matrix"
	"testing"
)

var n = 300

func Benchmark_MatrixMulGo(b *testing.B) {
	b.StopTimer()
	A := RandN(n, n)
	B := RandN(n, n)
	D := DiagFloat64(B.Array())
	D = D[:n]
	res := RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, D)
	}
}

func Benchmark_MatrixMulBlas(b *testing.B) {
	b.StopTimer()
	A := RandN(n, n)
	B := RandN(n, n)
	D := DiagFloat64(B.Array())
	D = D[:n]
	res := RandN(n, n)
	blas.Ops.Dgemm = cblas.Dgemm
	blas.Ops.Daxpy = cblas.Daxpy
	blas.Ops.Dcopy = cblas.Dcopy
	blas.Ops.Dscal = cblas.Dscal
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, D)
	}
} /*
func Benchmark_MatrixMulBlasGo(b *testing.B) {
	b.StopTimer()
	A := RandN(n, n)
	B := RandN(n, n)
	res := RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}*/

func Benchmark_Nrm2Go(b *testing.B) {
	A := RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_Nrm2Blas(b *testing.B) {
	A := RandN(n, n)
	v := A.VecView()
	blas.Ops.Dnrm2 = cblas.Dnrm2

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_DdotGo(b *testing.B) {
	A := RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		Ddot(v, v)
	}
}

func Benchmark_DdotBlas(b *testing.B) {
	A := RandN(n, n)
	v := A.VecView()
	blas.Ops.Ddot = cblas.Ddot

	for i := 0; i < b.N; i++ {
		Ddot(v, v)
	}
}

/*
func Benchmark_Nrm2GoFast(b *testing.B) {
	A := RandN(100, 100)
	v := A.VecView()
	blas.Dnrm2 = goblas.Dnrm2

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}*/

func TestMatrixBlas(t *testing.T) {
	A := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)
	B := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)

	res := NewDenseFloat64(2, 2)

	res.Mul(A, B)

	fmt.Println(res)

	A.T()
	B.T()
	res.Mul(A, B)

	fmt.Println(res)

	A.T()
	res.Mul(A, B)

	fmt.Println(res)
}
