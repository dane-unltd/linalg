package matrix

type Adder interface {
	Add(A, B Matrix)
}

type Muler interface {
	Add(A, B Matrix)
}

type Arith interface {
	Adder
	Muler
}

func Mul(A, B Expr) Expr {
	m, _ := A.Size()
	_, n := B.Size()
	return &mulExpr{m, n, A, B}
}
func Add(A, B Expr) Expr {
	m, _ := A.Size()
	_, n := B.Size()
	return &addExpr{m, n, A, B}
}
func Sub(A, B Expr) Expr {
	m, _ := A.Size()
	_, n := B.Size()
	return &subExpr{m, n, A, B}
}
