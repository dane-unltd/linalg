package matrix

type Expr interface {
	Size() (int, int)
	eval(res Arith, w *worker)
}

type mulExpr struct {
	m, n int
	A, B Expr
}

func (e *mulExpr) Size() (int, int) {
	return e.m, e.n
}

type divExpr struct {
	m, n int
	A, B Expr
}

func (e *divExpr) Size() (int, int) {
	return e.m, e.n
}

type addExpr struct {
	m, n int
	A, B Expr
}

func (e *addExpr) Size() (int, int) {
	return e.m, e.n
}

func (e *addExpr) eval(res Arith, w *worker) {
	Am, okA := e.A.(Matrix)
	Bm, okB := e.B.(Matrix)

	if !okA {
		Am = &w.stack[w.stackPos]
		w.stackPos++
		e.A.eval(Am, w)
	}
	if !okB {
		Bm = &w.stack[w.stackPos]
		w.stackPos++
		e.B.eval(Bm, w)
	}

	res.Add(Am, Bm)

	if !okA {
		w.stackPos--
	}

	if !okB {
		w.stackPos--
	}
}

type subExpr struct {
	m, n int
	A, B Expr
}

func (e *subExpr) Size() (int, int) {
	return e.m, e.n
}
