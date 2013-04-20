package goblas

import "github.com/kortschak/blas"

type Blas struct{}

var (
	_ blas.Float64 = Blas{}
)
