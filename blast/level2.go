package blast

func (A GeOp) MV(alpha float64, x Vector, beta float64, y Vector) {
	must(A.Check())
	must(y.Check())

	if A.Trans == NoTrans {
		if x.N != A.Cols || y.N != A.Rows {
			panic(ErrShape)
		}
		for i, j := 0, 0; i < A.Rows; i, j = i+1, j+y.Inc {
			r := A.Row(i)
			y.Data[j] = beta*y.Data[j] + alpha*Dot(r, x)
		}
		return
	}
	if A.Trans == Trans {
		if x.N != A.Rows || y.N != A.Cols {
			panic(ErrShape)
		}
		for i, j := 0, 0; i < A.Cols; i, j = i+1, j+y.Inc {
			c := A.Col(i)
			y.Data[j] = beta*y.Data[j] + alpha*Dot(c, x)
		}
		return
	}
	panic(ErrIllegalTrans)
}

func (T TrOp) Solve(alpha float64, x Vector) {
	must(T.Check())
	must(x.Check())

	panic("not implemented")
}
