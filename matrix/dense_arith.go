package matrix

import "fmt"

import "github.com/kortschak/blas"

func (res *Dense) Add(A, B Matrix) {
	m, n := res.Size()
	ma, na := A.Size()
	mb, nb := B.Size()

	if ma != m || na != n || mb != m || nb != n {
		fmt.Println(m, n, ma, na, mb, nb)
		panic("dimension missmatch")
	}
	switch A := A.(type) {
	case *Dense:
		switch B := B.(type) {
		case *Dense:
			for c := 0; c < n; c++ {
				res.ColView(c).Add(A.ColView(c), B.ColView(c))
			}
			return
		}
	}

	res.AddMM(A, B)
}

func (res *Dense) AddMM(A, B Matrix) {
	m, n := res.Size()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res.Set(i, j, A.At(i, j)+B.At(i, j))
		}
	}
}

func (res *Dense) Sub(A, B Matrix) {
	m, n := res.Size()
	ma, na := A.Size()
	mb, nb := B.Size()

	if ma != m || na != n || mb != m || nb != n {
		fmt.Println(m, n, ma, na, mb, nb)
		panic("dimension missmatch")
	}
	switch A := A.(type) {
	case *Dense:
		switch B := B.(type) {
		case *Dense:
			for c := 0; c < n; c++ {
				res.ColView(c).Sub(A.ColView(c), B.ColView(c))
			}
			return
		}
	}

	res.SubMM(A, B)
}

func (res *Dense) SubMM(A, B Matrix) {
	m, n := res.Size()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res.Set(i, j, A.At(i, j)-B.At(i, j))
		}
	}
}

func (res *Dense) Mul(A, B Matrix) {
	m, n := res.Size()
	ma, na := A.Size()
	mb, nb := B.Size()

	if ma != m || nb != n || na != mb {
		fmt.Println(m, n, ma, na, mb, nb)
		panic("dimension missmatch")
	}

	switch A := A.(type) {
	case *Dense:
		switch B := B.(type) {
		case *Dense:
			ops.Dgemm(blas.ColMajor, A.trans, B.trans, m, n, na,
				1, A.data, A.stride, B.data, B.stride, 0,
				res.data, res.stride)
			return
		case Vec:
			ops.Dgemm((blas.ColMajor), (A.trans), (blas.NoTrans),
				m, n, na, 1, A.data, A.stride, B, len(B), 0,
				res.data, res.stride)
			return
		case Diag:
			for i := 0; i < n; i++ {
				ops.Dcopy(m, A.ColView(i), 1, res.ColView(i), 1)
				ops.Dscal(m, B[i], res.ColView(i), 1)
			}
			return
		}
	case Diag:
		switch B := B.(type) {
		case Diag:
			for i := 0; i < m; i++ {
				res.Set(i, i, A[i]*B[i])
			}
			return
		}
	}

	res.MulMM(A, B)
}

//Slow general matrix multiplication
func (res *Dense) MulMM(A, B Matrix) {
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