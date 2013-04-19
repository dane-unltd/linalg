package linalg

import (
	"fmt"
	//	_ "github.com/dane-unltd/linalg/cblas"
	_ "github.com/dane-unltd/linalg/goblas"
	. "github.com/dane-unltd/linalg/matrix"
	"testing"
)

var n = 30

func Benchmark_MatrixMul(b *testing.B) {
	b.StopTimer()
	A := RandN(n, n)
	B := RandN(n, n)
	D := DiagFloat64(B.Array())
	D = D[:n]
	res := RandN(n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_Nrm2(b *testing.B) {
	A := RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		v.Nrm2()
	}
}

func Benchmark_Ddot(b *testing.B) {
	A := RandN(n, n)
	v := A.VecView()

	for i := 0; i < b.N; i++ {
		Ddot(v, v)
	}
}

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
