package linalg

import (
	"fmt"
	"github.com/dane-unltd/linalg/goblas"
	. "github.com/dane-unltd/linalg/matrix"
	"github.com/kortschak/cblas"
	"testing"
)

var n = 100

func Benchmark_MatrixMulCblas(b *testing.B) {
	RegisterOps(cblas.Blas{})
	b.StopTimer()
	A := RandN(n, n)
	B := RandN(n, n)
	D := Diag(B.Array())
	D = D[:n]
	res := RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_MatrixMulGo(b *testing.B) {
	RegisterOps(goblas.Blas{})
	b.StopTimer()
	A := RandN(n, n)
	B := RandN(n, n)
	D := Diag(B.Array())
	D = D[:n]
	res := RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_Nrm2Cblas(b *testing.B) {
	RegisterOps(cblas.Blas{})
	A := RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_Nrm2Go(b *testing.B) {
	RegisterOps(goblas.Blas{})
	A := RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_DdotCblas(b *testing.B) {
	RegisterOps(cblas.Blas{})
	A := RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		Ddot(v, v)
	}
}
func Benchmark_DdotGo(b *testing.B) {
	RegisterOps(goblas.Blas{})
	A := RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		Ddot(v, v)
	}
}

func TestMatrixBlas(t *testing.T) {
	RegisterOps(goblas.Blas{})
	A := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)
	B := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)

	res := NewDense(2, 2)

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
