package mat

import "github.com/dane-unltd/linalg/lapack"

func (D *Dense) Svd(S Vec, U, Vt *Dense, full bool) lapack.Info {
	if U.IsTr() || Vt.IsTr() {
		panic("cannot store factorization into transposed view")
	}

	if D.IsTr() {
		panic("factorizing transposed view of a matrix will lead to undesired results")
	}

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

	Dcp := make([]float64, D.stride*n)
	copy(Dcp, D.data)
	return ops.Dgesvd(jobu, jobvt, m, n, Dcp, D.stride, S,
		U.data, U.stride, Vt.data, Vt.stride)
}
