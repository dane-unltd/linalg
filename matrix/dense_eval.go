package matrix

func (D *Dense) FromExpr(e Expr) {
	done := make(chan struct{})
	req <- Request{D, e, done}

	<-done
}
