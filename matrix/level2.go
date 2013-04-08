package matrix

func (D *MatD) ApplyTo(v VecD) VecD {
	res := make(VecD, D.rows)
	Ops.Dgemv(D.IsTr(), D.rows, D.cols, 1, D.data, D.stride, v, 1, 0, res, 1)
	return res
}

func (D *MatD) ApplyToTr(v VecD) VecD {
	res := make(VecD, D.cols)
	Ops.Dgemv(!D.IsTr(), D.rows, D.cols, 1, D.data, D.stride, v, 1, 0, res, 1)
	return res
}
