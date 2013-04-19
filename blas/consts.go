package blas

type SrotmParams struct {
	Flag float32
	H    [4]float32 // Column-major 2 by 2 matrix.
}

type DrotmParams struct {
	Flag float64
	H    [4]float64 // Column-major 2 by 2 matrix.
}

type Order int

const (
	RowMajor Order = 101
	ColMajor Order = 102
)

type Transpose int

const (
	NoTrans   Transpose = 111
	Trans     Transpose = 112
	ConjTrans Transpose = 113
)

type Uplo int

const (
	Upper Uplo = 121
	Lower Uplo = 122
)

type Diag int

const (
	NonUnit Diag = 131
	Unit    Diag = 132
)

type Side int

const (
	Left  Side = 141
	Right Side = 142
)
