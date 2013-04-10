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
	A := RandN(100, 100)
	B := RandN(100, 100)
	res := RandN(100, 100)
	for i := 0; i < b.N; i++ {
		res.Mul(A, B)
	}
}

func Benchmark_BlasMul(b *testing.B) {
	A := RandN(100, 100)
	B := RandN(100, 100)
	res := RandN(100, 100)
	for i := 0; i < b.N; i++ {
		blas.DenseD{res}.Mul(A, B)
	}
}

func TestMatrixBlas(t *testing.T) {

	A := RandN(100, 100)
	B := RandN(100, 100)

	v := make(VecD, 100)
	D := make(DiagD, 100)

	resMat := NewDenseD(100, 100)
	resMat2 := NewDenseD(100, 1)

	blas.DenseD{resMat}.Mul(A, B)
	blas.DenseD{resMat}.Mul(A, D)
	blas.DenseD{resMat2}.Mul(A, v)

	resMat.Mul(A, B)
	resMat.Mul(A, D)
	resMat2.Mul(A, v)
}
