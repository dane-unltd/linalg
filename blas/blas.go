package blas

type MatD interface {
	Size() (int, int)
	Stride() int

	ArrayD() ([]float64, bool)
}
