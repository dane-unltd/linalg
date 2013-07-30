package mat

type LinOp interface {
	ApplyTo(v, dst Vec)
}

type Vectorer interface {
	Col(ix int, col Vec) Vec
}
