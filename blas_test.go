package linalg

import (
	"fmt"
	//_ "github.com/dane-unltd/linalg/blas"
	. "github.com/dane-unltd/linalg/matrix"
	_ "math"
	_ "math/rand"
	"testing"
)

func Benchmark_MatrixMul(b *testing.B) {
	A := RandN(100, 100)
	B := RandN(100, 100)
	res := RandN(100, 100)
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func TestMatrixBlas(t *testing.T) {
	A := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)

	D := DiagD{3, 4}

	res := NewDenseD(2, 2)

	res.Mul(A, D)

	fmt.Println(res)
}
