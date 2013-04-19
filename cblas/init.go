package cblas

import "github.com/dane-unltd/linalg/blas"

func init() {
	ops := BlasOps{}

	blas.Register(ops)
}
