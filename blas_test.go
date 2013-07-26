package linalg

import (
	"github.com/dane-unltd/engine/math3"
	"github.com/dane-unltd/linalg/clapack"
	"github.com/dane-unltd/linalg/goblas"
	"github.com/dane-unltd/linalg/mat"
	"github.com/kortschak/cblas"
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

/*
func Benchmark_MatrixMulEval(b *testing.B) {
	mat.Register(cblasops{})
	b.StopTimer()
	A := mat.RandN(n, n)
	B := mat.RandN(n, n)
	C := mat.RandN(n, n)
	D := mat.RandN(n, n)
	E := mat.RandN(n, n)
	res := mat.RandN(n, n)
	expr := mat.Mul(mat.Mul(mat.Mul(mat.Mul(A, B), C), D), E)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.FromExpr(expr)
	}
}*/

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

func TestMul(t *testing.T) {
	mat.Register(goblasops{})
	A := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	B := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	C := mat.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)

	temp := mat.NewDense(2)

	res := mat.NewDense(2)
	res1 := mat.NewFromArray([]float64{7, 10, 15, 22}, false, 2, 2)

	res.Mul(A, B)

	temp.Sub(res, res1)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, res1)
	}

	temp.Mul(res, C)

	At := A.TrView()
	Bt := B.TrView()
	res.Mul(At, Bt)

	res.Mul(At, B)

}
