package matrix

import "fmt"

import "github.com/dane-unltd/linalg/blas"

func (res *DenseD) Add(A, B Matrix) {
	m, n := res.Size()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res.Set(i, j, A.At(i, j)+B.At(i, j))
		}
	}
}

func (res *DenseD) Sub(A, B Matrix) {
	m, n := res.Size()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res.Set(i, j, A.At(i, j)-B.At(i, j))
		}
	}
}

func (res *DenseD) Mul(A, B Matrix) {
	m, n := res.Size()
	ma, na := A.Size()
	mb, nb := B.Size()

	if ma != m || nb != n || na != mb {
		fmt.Println(m, n, ma, na, mb, nb)
		panic("dimension missmatch")
	}

	switch A := A.(type) {
	case *DenseD:
		switch B := B.(type) {
		case *DenseD:
			blas.Dgemm(int(blas.ColMajor), int(A.trans), int(B.trans), m, n, na,
				1, A.data, A.stride, B.data, B.stride, 0,
				res.data, res.stride)
			return
		case VecD:
			blas.Dgemm(int(blas.ColMajor), int(A.trans), int(blas.NoTrans),
				m, n, na, 1, A.data, A.stride, B, len(B), 0,
				res.data, res.stride)
			return
		}
	}

	res.MulMM(A, B)
}

//Slow general matrix multiplication
func (res *DenseD) MulMM(A, B Matrix) {
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
