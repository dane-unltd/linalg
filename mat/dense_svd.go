package mat

import "github.com/dane-unltd/linalg/lapack"

func (D *Dense) Svd(S Vec, U, Vt *Dense, full bool) lapack.Info {
	if U.view || Vt.view {
		panic("cannot use view of a matrix as svd input")
	}

	m, n := D.Size()

	nSV := m
	if n < m {
		nSV = n
	}

	var jobu, jobvt lapack.Job

	if full {
		jobu = 'A'
		jobvt = 'A'
		U.rows = m
		U.cols = m
		U.stride = m
		Vt.rows = n
		Vt.cols = n
		Vt.stride = n
	} else {
		jobu = 'S'
		jobvt = 'S'
		U.rows = m
		U.cols = nSV
		U.stride = m
		Vt.rows = nSV
		Vt.cols = n
		Vt.stride = nSV
	}
	if len(Vt.data) < Vt.rows*Vt.cols {
		Vt.data = make([]float64, Vt.rows*Vt.cols)
	}
	if len(U.data) < U.rows*U.cols {
		U.data = make([]float64, U.rows*U.cols)
	}

	Dcp := make([]float64, D.stride*n)
	copy(Dcp, D.data)
	return ops.Dgesvd(jobu, jobvt, m, n, Dcp, D.Stride(), S,
		U.Array(), U.Stride(), Vt.Array(), Vt.Stride())
}
