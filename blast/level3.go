package blast

func (A Ge) Norm(norm Norm) float64 {
	must(A.Check())
	switch norm {
	case One, RealOne:
	case Two:
	case Frobenius:
	case Inf, RealInf:
	case Max, RealMax:
	}

	return 0
}

func (A Ge) DiagScale(side Side, d Vector) {
	must(A.Check())
	must(d.Check())

	if side == Left {
		if d.N != A.Rows {
			panic(ErrShape)
		}
		for i, j := 0, 0; i < d.N*d.Inc; i, j = i+d.Inc, j+1 {
			r := A.Row(j)
			Axpby(0, Vector{}, d.Data[i], r)
		}
		return
	}

	if side == Right {
		if d.N != A.Cols {
			panic(ErrShape)
		}
		for i, j := 0, 0; i < d.N*d.Inc; i, j = i+d.Inc, j+1 {
			c := A.Col(j)
			Axpby(0, Vector{}, d.Data[i], c)
		}
		return
	}
	panic(ErrIllegalSide)
}

func (A Ge) LRScale(dl, dr Vector) {
	A.DiagScale(Left, dl)
	A.DiagScale(Right, dr)
}

func (B Ge) GeAcc(alpha float64, A GeOp, beta float64) {
	must(A.Check())
	must(B.Check())
	if A.Trans == NoTrans {
		if A.Rows != B.Rows || A.Cols != B.Cols {
			panic(ErrShape)
		}
		for i := 0; i < A.Cols; i++ {
			ca, cb := A.Col(i), B.Col(i)
			Axpby(alpha, ca, beta, cb)
		}
	}
	if A.Trans == Trans {
		if A.Cols != B.Rows || A.Rows != B.Cols {
			panic(ErrShape)
		}
		for i := 0; i < A.Cols; i++ {
			ca, rb := A.Col(i), B.Row(i)
			Axpby(alpha, ca, beta, rb)
		}
	}
	panic(ErrIllegalTrans)
}

func (C Ge) GeAdd(alpha float64, A Ge, beta float64, B Ge) {
	must(A.Check())
	must(B.Check())
	must(C.Check())

	if A.Rows != B.Rows || A.Cols != B.Cols || C.Rows != A.Rows ||
		C.Cols != A.Cols {
		panic(ErrShape)
	}
	for i := 0; i < A.Cols; i++ {
		ca, cb, cc := A.Col(i), B.Col(i), C.Col(i)
		Waxpby(alpha, ca, beta, cb, cc)
	}
}

func (A GeOp) MM(alpha float64, B GeOp, beta float64, C Ge) {
	must(A.Check())
	must(B.Check())
	must(C.Check())

	if A.Trans == NoTrans {
		if A.Rows != C.Rows || (B.Trans == NoTrans && A.Cols != B.Rows) ||
			(B.Trans == Trans && A.Cols != B.Cols) {
			panic(ErrShape)
		}
	} else if A.Trans == Trans {
		if A.Cols != C.Rows || (B.Trans == NoTrans && A.Rows != B.Rows) ||
			(B.Trans == Trans && A.Rows != B.Cols) {
			panic(ErrShape)
		}
	} else {
		panic(ErrIllegalTrans)
	}

	if B.Trans == NoTrans {
		if B.Cols != C.Cols {
			panic(ErrShape)
		}
		for i := 0; i < C.Cols; i++ {
			cb, cc := B.Col(i), C.Col(i)
			A.MV(alpha, cb, beta, cc)
		}

	} else if B.Trans == Trans {
		if B.Rows != C.Cols {
			panic(ErrShape)
		}
		for i := 0; i < C.Cols; i++ {
			rb, cc := B.Row(i), C.Col(i)
			A.MV(alpha, rb, beta, cc)
		}
	} else {
		panic(ErrIllegalTrans)
	}

}
