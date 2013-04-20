package matrix

import "github.com/kortschak/blas"

var ops matops

type matops interface {
	blas.Float64
}

func RegisterOps(o matops) {
	ops = o
}
