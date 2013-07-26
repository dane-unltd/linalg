package mat

/*
type Request struct {
	res  Arith
	expr Expr
	done chan<- struct{}
}

type worker struct {
	stack    [50]Dense
	stackPos int
}

var req = make(chan Request)

func (w *worker) run() {
	for {
		r := <-req
		r.expr.eval(r.res, w)
		r.done <- struct{}{}
	}
}

var wrk = &worker{}

func init() {
	go wrk.run()
}*/
