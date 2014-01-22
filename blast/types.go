package blast

// Type Transpose is used to specify the transposition operation for a
// routine.
type Transpose int

const (
	NoTrans Transpose = 111 + iota
	Trans
	ConjTrans
)

// Type Uplo is used to specify whether the matrix is an upper or lower
// triangular matrix.
type Uplo int

const (
	Upper Uplo = 121 + iota
	Lower
)

// Type Diag is used to specify whether the matrix is a unit or non-unit
// triangular matrix.
type Diag int

const (
	NonUnit Diag = 131 + iota
	Unit
)

// Type side is used to specify from which side a multiplication operation
// is performed.
type Side int

const (
	Left Side = 141 + iota
	Right
)

// Type norm is used to specify which norm to calculate
type Norm int

const (
	One Norm = 171 + iota
	RealOne
	Two
	Frobenius
	Inf
	RealInf
	Max
	RealMax
)

type Ge struct {
	Rows, Cols int
	Data       []float64
	Stride     int
}

type GeOp struct {
	Ge
	Trans Transpose
}

type Tr struct {
	Ge
	Uplo Uplo
	Diag Diag
}

type TrOp struct {
	Tr
	Trans Trans
	Side  side
}

type Sy struct {
	Ge
	Uplo Uplo
}

type SyOp struct {
	Sy
	Side Side
}

func (A Ge) Check() error {
	if A.Stride <= 1 || A.Stride <= A.Rows {
		return ErrIllegalStride
	}
	if len(A.Data) < A.Stride*(A.Cols-1)+A.Rows {
		return ErrNotEnoughData
	}
	return nil
}

func (A Ge) Row(i int) Vector {
	return Vector{A.Cols, A.Data[i:], A.Stride}
}

func (A Ge) Col(i int) Vector {
	return Vector{A.Rows, A.Data[i*A.Stride:], 1}
}

func (A Ge) Sub(i, j, r, c int) Ge {
	B = Ge{r, c, A.Data[i+A.Stride*j:], A.Stride}
	must(B.Check())
	return B
}

type Vector struct {
	N    int
	Data []float64
	Inc  int
}

func (v Vector) Check() error {
	if v.Inc == 0 {
		return ErrZeroInc
	}
	if len(v.Data) < v.Inc*(v.N-1)+1 {
		return ErrNotEnoughData
	}
	return nil
}

func Slice2Vec(d []float64) {
	return Vector{len(d), d, 1}
}

type Perm []int

type Error string

func (err Error) Error() string { return string(err) }

const (
	ErrShape         = Error("blast: dimension mismatch")
	ErrIllegalStride = Error("blast: illegal stride")
	ErrIllegalTrans  = Error("blast: illegal transpose")
	ErrIllegalSide   = Error("blast: illegal side")
	ErrNotEnoughData = Error("blast: data slice to short")
	ErrZeroInc       = Error("blas: Inc is 0")
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}
