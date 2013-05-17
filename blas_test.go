package linalg

import (
	"github.com/dane-unltd/linalg/clapack"
	"github.com/dane-unltd/linalg/goblas"
	"github.com/dane-unltd/linalg/math3"
	"github.com/dane-unltd/linalg/matrix"
	"github.com/kortschak/cblas"
	"testing"
)

var n = 50

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
	matrix.Register(cblasops{})
	b.StopTimer()
	A := matrix.RandN(n, n)
	B := matrix.RandN(n, n)
	C := matrix.RandN(n, n)
	D := matrix.RandN(n, n)
	E := matrix.RandN(n, n)
	res1 := matrix.RandN(n, n)
	res2 := matrix.RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res1.Mul(A, B)
		res2.Mul(res1, C)
		res1.Mul(res2, D)
		res2.Mul(res1, E)
	}
}

func Benchmark_MatrixMulGo(b *testing.B) {
	matrix.Register(goblasops{})
	b.StopTimer()
	A := matrix.RandN(n, n)
	B := matrix.RandN(n, n)
	C := matrix.RandN(n, n)
	D := matrix.RandN(n, n)
	E := matrix.RandN(n, n)
	res1 := matrix.RandN(n, n)
	res2 := matrix.RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res1.Mul(A, B)
		res2.Mul(res1, C)
		res1.Mul(res2, D)
		res2.Mul(res1, E)
	}
}

func Benchmark_MatrixMulEval(b *testing.B) {
	matrix.Register(goblasops{})
	b.StopTimer()
	A := matrix.RandN(n, n)
	B := matrix.RandN(n, n)
	C := matrix.RandN(n, n)
	D := matrix.RandN(n, n)
	E := matrix.RandN(n, n)
	res := matrix.RandN(n, n)
	expr := matrix.Mul(matrix.Mul(matrix.Mul(matrix.Mul(A, B), C), D), E)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.FromExpr(expr)
	}
}

func Benchmark_Nrm2Cblas(b *testing.B) {
	matrix.Register(cblasops{})
	A := matrix.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_Nrm2Go(b *testing.B) {
	matrix.Register(goblasops{})
	A := matrix.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_DdotCblas(b *testing.B) {
	matrix.Register(cblasops{})
	A := matrix.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		matrix.Dot(v, v)
	}
}

func Benchmark_DdotGo(b *testing.B) {
	matrix.Register(goblasops{})
	A := matrix.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		matrix.Dot(v, v)
	}
}

func TestMul(t *testing.T) {
	matrix.Register(goblasops{})
	A := matrix.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	B := matrix.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)
	C := matrix.NewFromArray([]float64{1, 2, 3, 4}, false, 2, 2)

	mul := matrix.Mul(matrix.Mul(A, B), C)
	resE := matrix.NewDense(2)
	resE.FromExpr(mul)

	temp := matrix.NewDense(2)

	res := matrix.NewDense(2)
	res1 := matrix.NewFromArray([]float64{7, 10, 15, 22}, false, 2, 2)

	res.Mul(A, B)

	temp.Sub(res, res1)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", res, res1)
	}

	temp.Mul(res, C)
	temp.Sub(temp, resE)
	if temp.VecView().Nrm2Sq() > 0.01 {
		t.Error("wrong result", temp)
	}

	A.T()
	B.T()
	res.Mul(A, B)

	A.T()
	res.Mul(A, B)

}
