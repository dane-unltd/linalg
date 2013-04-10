package lapack

import "github.com/dane-unltd/linalg/matrix"

func DenseDSvd(D *matrix.DenseD, S matrix.DiagD, U, Vt *matrix.DenseD) {
	if U.IsTr() {
		U.Tr()
	}
	if Vt.IsTr() {
		Vt.Tr()
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

	/*info := */ Dgesvd(jobu, jobvt, m, n, D.Copy().(*matrix.DenseD).ArrayD(), D.Stride(), S,
		U.ArrayD(), U.Stride(), Vt.ArrayD(), Vt.Stride())
}
