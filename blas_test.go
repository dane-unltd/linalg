package linalg

import (
	_ "fmt"
	"github.com/dane-unltd/linalg/blas"
	. "github.com/dane-unltd/linalg/matrix"
	_ "math"
	_ "math/rand"
	"testing"
)

func Benchmark_MatrixMul(b *testing.B) {
	A := RandN(1000, 1000)
	B := RandN(1000, 1000)
	res := RandN(1000, 1000)
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_BlasMul(b *testing.B) {
	A := RandN(1000, 1000)
	B := RandN(1000, 1000)
	res := blas.DenseD{RandN(1000, 1000)}
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func TestMatrixBlas(t *testing.T) {

	A := RandN(100, 100)
	B := RandN(100, 100)

	v := make(VecD, 100)
	D := make(DiagD, 100)

	resBlas := blas.DenseD{NewDenseD(100, 100)}
	resMat := NewDenseD(100, 100)
	resBlas2 := blas.DenseD{NewDenseD(100, 1)}
	resMat2 := NewDenseD(100, 1)

	resBlas.Mul(A, B)
	resBlas.Mul(A, D)
	resBlas2.Mul(A, v)

	resMat.Mul(A, B)
	resMat.Mul(A, D)
	resMat2.Mul(A, v)
}
