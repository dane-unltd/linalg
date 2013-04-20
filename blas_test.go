package linalg

import (
	"fmt"
	"github.com/dane-unltd/linalg/goblas"
	"github.com/dane-unltd/linalg/matrix"
	"github.com/kortschak/cblas"
	"testing"
)

var n = 100

func Benchmark_MatrixMulCblas(b *testing.B) {
	matrix.Register(cblas.Blas{})
	b.StopTimer()
	A := matrix.RandN(n, n)
	B := matrix.RandN(n, n)
	D := matrix.Diag(B.Array())
	D = D[:n]
	res := matrix.RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_MatrixMulGo(b *testing.B) {
	matrix.Register(goblas.Blas{})
	b.StopTimer()
	A := matrix.RandN(n, n)
	B := matrix.RandN(n, n)
	D := matrix.Diag(B.Array())
	D = D[:n]
	res := matrix.RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_Nrm2Cblas(b *testing.B) {
	matrix.Register(cblas.Blas{})
	A := matrix.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_Nrm2Go(b *testing.B) {
	matrix.Register(goblas.Blas{})
	A := matrix.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_DdotCblas(b *testing.B) {
	matrix.Register(cblas.Blas{})
	A := matrix.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		matrix.Dot(v, v)
	}
}
func Benchmark_DdotGo(b *testing.B) {
	matrix.Register(goblas.Blas{})
	A := matrix.RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		matrix.Dot(v, v)
	}
}

func TestMatrixBlas(t *testing.T) {
	matrix.Register(goblas.Blas{})
	A := matrix.FromArray([]float64{1, 2, 3, 4}, true, 2, 2)
	B := matrix.FromArray([]float64{1, 2, 3, 4}, true, 2, 2)

	res := matrix.NewDense(2, 2)

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
