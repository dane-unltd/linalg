package lapack

import (
	//"errors"
	"github.com/dane-unltd/linalg/dense"
)

func SvdD(A MatD, full ...bool) (S, U, Vt *dense.MatD) {
	var f bool
	if len(full) > 0 {
		f = full[0]
	}
	m, n := A.Size()
	stride := A.Stride()
	Aa, _ := A.ArrayD()

	nSV := min(m, n)
	if f {
		U = dense.Zeros(m)
		Sa := make([]float64, nSV)
		Vt = dense.Zeros(n)

		Ua, _ := U.ArrayD()
		Vta, _ := Vt.ArrayD()

		info := dgesvd("A", "A", m, n, Aa, stride,
			Sa, Ua, m, Vta, n)
		if info != 0 {
			panic("lapack error")
		}

		S = dense.Zeros(m, n)
		for i := 0; i < nSV; i++ {
			S.Set(Sa[i], i, i)
		}
	} else {
		U = dense.Zeros(m, nSV)
		Sa := make([]float64, nSV)
		Vt = dense.Zeros(nSV, n)

		Ua, _ := U.ArrayD()
		Vta, _ := Vt.ArrayD()

		info := dgesvd("S", "S", m, n, Aa, stride,
			Sa, Ua, m, Vta, nSV)
		if info != 0 {
			panic("lapack error")
		}

		S = dense.Diag(Sa)
	}
	return
}
