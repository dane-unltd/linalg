package matrix

func (R *MatD) SolveTriU(v VecD) VecD {
	res := v.Copy()
	n := len(v)
	Ops.Dtrsv("U", R.IsTr(), "N", n, R.data, R.stride, res, 1)
	return res
}

func (R *MatD) SolveTriL(v VecD) VecD {
	res := v.Copy()
	n := len(v)
	Ops.Dtrsv("L", R.IsTr(), "N", n, R.data, R.stride, res, 1)
	return res
}

func Lsqr(A *MatD, b VecD) VecD {
	m, n := A.Size()

	Atr := A.Copy()
	Atr.Tr()

	if m > n {
		AtA := MulD(Atr, A)

		L := AtA.Chol()
		res := L.SolveTriL(Atr.ApplyTo(b))
		L.Tr()
		res = L.SolveTriL(res)
		return res
	} else if n > m {
		return nil
	} else {
		return nil
	}
	return nil
}
