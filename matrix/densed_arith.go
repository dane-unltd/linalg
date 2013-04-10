package matrix

func (res *DenseD) Add(A, B Matrix) {
	m, n := res.Size()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res.Set(i, j, A.At(i, j)+B.At(i, j))
		}
	}
}

func (res *DenseD) Mul(A, B Matrix) {
	Mul(res, A, B)
}

func (res *DenseD) MulMDiag(A Matrix, D DiagD) {
	println("matrix diag")
	for i := 0; i < res.rows; i++ {
		for j := 0; j < res.cols; j++ {
			res.Set(i, j, A.At(i, j)*D[i])
		}
	}
}

func (res *DenseD) MulMVec(A Matrix, v VecD) {
	println("matrix vec")
	for i := 0; i < res.rows; i++ {
		res.data[i] = 0
		for k := 0; k < len(v); k++ {
			res.data[i] += A.At(i, k) * v[k]
		}
	}
}

func (res *DenseD) MulMM(A, B Matrix) {
	println("matrix.MulMM")
	_, K := A.Size()
	for i := 0; i < res.rows; i++ {
		for j := 0; j < res.cols; j++ {
			res.data[res.dataIx(i, j)] = 0
			for k := 0; k < K; k++ {
				res.data[res.dataIx(i, j)] += A.At(i, k) *
					B.At(k, j)
			}
		}
	}
}

func (res *DenseD) MulDD(A, B *DenseD) {
	println("matrix.MulDD")
	res.MulMM(A, B)
}
