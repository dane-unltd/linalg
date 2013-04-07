package matrix

func (S *MatD) Chol() *MatD {
	Sc := S.Copy()
	Ops.Dpotrf("L", Sc.rows, Sc.data, Sc.stride)
	return Sc
}

func (S *MatD) EigSy() (V *MatD, d VecD) {
	n := S.rows
	d = make(VecD, n)
	V = S.Copy()
	Ops.Dsyevd("V", "L", n, V.data, V.stride, d)
	return
}
