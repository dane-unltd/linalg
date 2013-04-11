package blas

import "github.com/dane-unltd/linalg/matrix"

func init() {
	matrix.VecDNrm2 = VecDNrm2
}

func VecDNrm2(v matrix.VecD) float64 {
	return Dnrm2(len(v), v, 1)
}
