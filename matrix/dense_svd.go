package matrix

import "github.com/dane-unltd/linalg/lapack"

func (D *Dense) Svd(S Diag, U, Vt *Dense) lapack.Info {
	if U.IsTr() {
		U.T()
	}
	if Vt.IsTr() {
		Vt.T()
	}

	m, n := D.Size()
	mu, nu := U.Size()
	mv, nv := Vt.Size()

	if mu != m || nv != n || nu != len(S) || mv != len(S) {
		panic("dimension missmatch")
	}

	nSV := m
	if n < m {
		nSV = n
	}

	var jobu, jobvt lapack.Job

	if nu == mu {
		jobu = 'A'
	} else if nu == nSV {
		jobu = 'S'
	} else {
		panic("wrong dimensions of U")
	}

	if nv == mv {
		jobvt = 'A'
	} else if mv == nSV {
		jobvt = 'S'
	} else {
		panic("wrong dimensions of V")
	}

	Dcp := make([]float64, D.stride*n)
	copy(Dcp, D.data)
	return ops.Dgesvd(jobu, jobvt, m, n, Dcp, D.Stride(), S,
		U.Array(), U.Stride(), Vt.Array(), Vt.Stride())
}
