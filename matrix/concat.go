package matrix

func ConcatD(A, B *MatD) *MatD {
	if A.rows != B.rows {
		panic("dimension missmatch")
	}
	res := Zeros(A.rows, A.cols+B.cols)

	col := 0
	for i := 0; i < A.cols; i++ {
		copy(res.data[col*res.stride:],
			A.data[i*A.stride:(i*A.stride+A.rows)])
		col++
	}
	for i := 0; i < B.cols; i++ {
		copy(res.data[col*res.stride:],
			B.data[i*B.stride:(i*B.stride+B.rows)])
		col++
	}
	return res
}
