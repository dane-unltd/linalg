package matrix

import (
//"errors"
)

func SvdD(A MatDable, full ...bool) (S, U, Vt *MatD) {
	var f bool
	if len(full) > 0 {
		f = full[0]
	}
	m, n := A.Size()
	stride := A.Stride()
	Aa := A.ArrayD()

	nSV := n
	if m < n {
		nSV = m
	}

	if f {
		U = Zeros(m)
		Sa := make([]float64, nSV)
		Vt = Zeros(n)

		Ua := U.ArrayD()
		Vta := Vt.ArrayD()

		info := Ops.Dgesvd("A", "A", m, n, Aa, stride,
			Sa, Ua, m, Vta, n)
		if info != 0 {
			panic("lapack error")
		}

		S = Zeros(m, n)
		for i := 0; i < nSV; i++ {
			S.Set(Sa[i], i, i)
		}
	} else {
		U = Zeros(m, nSV)
		Sa := make([]float64, nSV)
		Vt = Zeros(nSV, n)

		Ua := U.ArrayD()
		Vta := Vt.ArrayD()

		info := Ops.Dgesvd("S", "S", m, n, Aa, stride,
			Sa, Ua, m, Vta, nSV)
		if info != 0 {
			panic("lapack error")
		}

		S = Diag(Sa)
	}
	if A.IsTr() {
		temp := U
		U = Vt
		Vt = temp
		U.Tr()
		Vt.Tr()
	}
	return
}
