package mat

import "github.com/dane-unltd/linalg/lapack"
import "github.com/gonum/blas"

func (D *Dense) Svd(S Vec, U, Vt *Dense, full bool) lapack.Info {
	m, n := D.Dims()

	nSV := m
	if n < m {
		nSV = n
	}

	var jobu, jobvt lapack.Job

	if full {
		jobu = 'A'
		jobvt = 'A'
		U.recvDimCheck(m, m)
		Vt.recvDimCheck(n, n)
	} else {
		jobu = 'S'
		jobvt = 'S'
		U.recvDimCheck(m, nSV)
		Vt.recvDimCheck(nSV, n)
	}

	Dcp := make([]float64, D.rows*n)
	superb := make([]float64, nSV)
	copy(Dcp, D.data)
	return ops.Dgesvd(blas.ColMajor, jobu, jobvt, m, n, Dcp, D.rows, S,
		U.data, U.rows, Vt.data, Vt.rows, superb)
}
