package mat

/*
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

func (e *mulExpr) eval(res Arith, w *worker) {
	Am, okA := e.A.(Matrix)
	Bm, okB := e.B.(Matrix)

	if !okA {
		D := &w.stack[w.stackPos]
		w.stackPos++

		m, n := e.A.Size()
		if len(D.data) < m*n {
			D.data = make([]float64, m*n)
		}
		e.A.eval(D, w)
		Am = D
	}
	if !okB {
		D := &w.stack[w.stackPos]
		w.stackPos++

		m, n := e.B.Size()
		if len(D.data) < m*n {
			D.data = make([]float64, m*n)
		}
		e.B.eval(D, w)
		Bm = D
	}

	res.Mul(Am, Bm)

	if !okA {
		w.stackPos--
	}

	if !okB {
		w.stackPos--
	}
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
		D := &w.stack[w.stackPos]
		w.stackPos++

		m, n := e.A.Size()
		if len(D.data) < m*n {
			D.data = make([]float64, m*n)
		}
		e.A.eval(D, w)
		Am = D
	}
	if !okB {
		D := &w.stack[w.stackPos]
		w.stackPos++

		m, n := e.B.Size()
		if len(D.data) < m*n {
			D.data = make([]float64, m*n)
		}
		e.B.eval(D, w)
		Bm = D
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

func (e *subExpr) eval(res Arith, w *worker) {
	Am, okA := e.A.(Matrix)
	Bm, okB := e.B.(Matrix)

	if !okA {
		D := &w.stack[w.stackPos]
		w.stackPos++

		m, n := e.A.Size()
		if len(D.data) < m*n {
			D.data = make([]float64, m*n)
		}
		e.A.eval(D, w)
		Am = D
	}
	if !okB {
		D := &w.stack[w.stackPos]
		w.stackPos++

		m, n := e.B.Size()
		if len(D.data) < m*n {
			D.data = make([]float64, m*n)
		}
		e.B.eval(D, w)
		Bm = D
	}

	res.Sub(Am, Bm)

	if !okA {
		w.stackPos--
	}

	if !okB {
		w.stackPos--
	}
}
*/
