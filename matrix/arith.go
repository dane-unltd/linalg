package matrix

type Expr interface {
	Size() (int, int)
}

type mulExpr struct {
	m, n int
	A, B Expr
}

func (e *mulExpr) Size() (int, int) {
	return m, n
}

type divExpr struct {
	m, n int
	A, B Expr
}

func (e *divExpr) Size() (int, int) {
	return m, n
}

type addExpr struct {
	m, n int
	A, B Expr
}

func (e *addExpr) Size() (int, int) {
	return m, n
}

type subExpr struct {
	m, n int
	A, B Expr
}

func (e *subExpr) Size() (int, int) {
	return m, n
}
