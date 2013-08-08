package linalg

import (
	"github.com/dane-unltd/engine/math3"
	"github.com/dane-unltd/linalg/clapack"
	"github.com/dane-unltd/linalg/goblas"
	"github.com/dane-unltd/linalg/mat"
	"github.com/gonum/blas/cblas"
	"testing"
)

var n = 3

type cblasops struct {
	cblas.Blas
	clapack.Lapack
}

type goblasops struct {
	goblas.Blas
	clapack.Lapack
}

func Benchmark_MatrixMulMath3(b *testing.B) {
	b.StopTimer()
	A := math3.Mat{1, 2, 3, 4, 5, 6, 7, 8, 9}
	B := math3.Mat{1, 2, 3, 4, 5, 6, 7, 8, 9}
	C := math3.Mat{1, 2, 3, 4, 5, 6, 7, 8, 9}
	D := math3.Mat{1, 2, 3, 4, 5, 6, 7, 8, 9}
	E := math3.Mat{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res1 := math3.Mat{}
	res2 := math3.Mat{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res1.Mul(&A, &B)
		res2.Mul(&res1, &C)
		res1.Mul(&res2, &D)
		res2.Mul(&res1, &E)
	}
}

func Benchmark_MatrixMulCblas(b *testing.B) {
	mat.Register(cblasops{})
	b.StopTimer()
	A := mat.RandN(n, n)
	B := mat.RandN(n, n)
	C := mat.RandN(n, n)
	D := mat.RandN(n, n)
	E := mat.RandN(n, n)
	res1 := mat.RandN(n, n)
	res2 := mat.RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res1.Mul(A, B)
		res2.Mul(res1, C)
		res1.Mul(res2, D)
		res2.Mul(res1, E)
	}
}

func Benchmark_MatrixMulGo(b *testing.B) {
	mat.Register(goblasops{})
	b.StopTimer()
	A := mat.RandN(n, n)
	B := mat.RandN(n, n)
	C := mat.RandN(n, n)
	D := mat.RandN(n, n)
	E := mat.RandN(n, n)
	res1 := mat.RandN(n, n)
	res2 := mat.RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res1.Mul(A, B)
		res2.Mul(res1, C)
		res1.Mul(res2, D)
		res2.Mul(res1, E)
	}
}

func Benchmark_Nrm2Cblas(b *testing.B) {
	mat.Register(cblasops{})
	A := mat.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_Nrm2Go(b *testing.B) {
	mat.Register(goblasops{})
	A := mat.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_DdotCblas(b *testing.B) {
	mat.Register(cblasops{})
	A := mat.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		mat.Dot(v, v)
	}
}

func Benchmark_DdotGo(b *testing.B) {
	mat.Register(goblasops{})
	A := mat.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		mat.Dot(v, v)
	}
}
