package lapack

import "github.com/dane-unltd/linalg/matrix"

func Dgsvd(D *matrix.Dense, S matrix.Diag, U, Vt *matrix.Dense) {
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

	jobu, jobvt := "", ""

	if nu == mu {
		jobu = "A"
	} else if nu == nSV {
		jobu = "S"
	} else {
		panic("wrong dimensions of U")
	}

	if nv == mv {
		jobvt = "A"
	} else if mv == nSV {
		jobvt = "S"
	} else {
		panic("wrong dimensions of V")
	}

	/*info := */ Dgesvd(jobu, jobvt, m, n, D.Copy().(*matrix.Dense).Array(), D.Stride(), S,
		U.Array(), U.Stride(), Vt.Array(), Vt.Stride())
}
