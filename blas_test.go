package linalg

import (
	"fmt"
	_ "github.com/dane-unltd/linalg/blasinit"
	. "github.com/dane-unltd/linalg/matrix"
	"math"
	"math/rand"
	"testing"
)

func Benchmark_MulDelayed(b *testing.B) {
	A := RandN(1000, 1000)
	B := RandN(1000, 1000)
	C := RandN(1000, 1)
	for i := 0; i < b.N; i++ {
		MulD(A, B, C)
	}
}

func Benchmark_MulInstant(b *testing.B) {
	A := RandN(1000, 1000)
	B := RandN(1000, 1000)
	C := RandN(1000, 1)
	for i := 0; i < b.N; i++ {
		MulD(MulD(A, B), C)
	}
}

func Benchmark_Add(b *testing.B) {
	x := ZeroVec(10000)
	for i := range x {
		x[i] = rand.NormFloat64()
	}
	y := ZeroVec(10000)
	for i := range y {
		y[i] = rand.NormFloat64()
	}

	res := ZeroVec(10000)

	for i := 0; i < b.N; i++ {
		res.Add(x, y)
	}
}
func Benchmark_AddaY(b *testing.B) {
	x := ZeroVec(10000)
	for i := range x {
		x[i] = rand.NormFloat64()
	}
	y := ZeroVec(10000)
	for i := range y {
		y[i] = rand.NormFloat64()
	}

	for i := 0; i < b.N; i++ {
		res := x.Copy()
		res.AddaY(y)
	}
}

func TestAdd(t *testing.T) {
	x := ZeroVec(10000)
	for i := range x {
		x[i] = rand.NormFloat64()
	}
	y := ZeroVec(10000)
	for i := range y {
		y[i] = rand.NormFloat64()
	}

	res := ZeroVec(10000)

	res.Add(x, y)
	res2 := x.Copy()
	res2.AddaY(y)
	if math.Abs(res.Sub(res, res2).Nrm2()) > 0.0001 {
		t.Error("MulD failed", res.Nrm2())
	}
}

func Test_MulD(t *testing.T) {
	A := FromArrayD([]float64{1, 2}, true, 1, 2)
	B := FromArrayD([]float64{1, 2, 3, 4}, true, 2, 2)

	fmt.Println(MulD(A, B))
}
