package matrix

import "reflect"

var DenseDMul = make(map[TypePair]func(*DenseD, Matrix, Matrix))

//registering fast paths
func init() {
	tp := TypePair{reflect.TypeOf((*DenseD)(nil)), reflect.TypeOf(DiagD(nil))}
	DenseDMul[tp] = DenseDMulAD
	tp = TypePair{reflect.TypeOf((*DenseD)(nil)), reflect.TypeOf(VecD(nil))}
	DenseDMul[tp] = DenseDMulAv
	tp = TypePair{reflect.TypeOf((*DenseD)(nil)), reflect.TypeOf((*DenseD)(nil))}
	DenseDMul[tp] = DenseDMulAA
}

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
		panic("dimension missmatch")
	}

	tp := TypePair{reflect.TypeOf(A), reflect.TypeOf(B)}

	op, ok := DenseDMul[tp]

	if ok {
		op(res, A, B)
	} else {
		res.MulMM(A, B)
	}
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

//Faster specialized multiplications
func DenseDMulAD(res *DenseD, A Matrix, B Matrix) {
	D := B.(DiagD)
	for i := 0; i < res.rows; i++ {
		for j := 0; j < res.cols; j++ {
			res.Set(i, j, A.At(i, j)*D[j])
		}
	}
}

func DenseDMulAv(res *DenseD, A Matrix, B Matrix) {
	v := B.(VecD)
	for i := 0; i < res.rows; i++ {
		res.data[i] = 0
		for k := 0; k < len(v); k++ {
			res.data[i] += A.At(i, k) * v[k]
		}
	}
}

func DenseDMulAA(res *DenseD, A, B Matrix) {
	res.MulMM(A, B)
}
