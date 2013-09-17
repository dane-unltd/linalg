package mat

type Operator interface {
	ApplyTo(v, dst Vec)
	ApplyInverseTo(v, dst Vec)
	Dims() (int, int)
}

type Vectorer interface {
	Col(ix int, col Vec) Vec
}

type Viewer interface {
	View(i, j, r, c int) *denseView
}

type Coler interface {
	Col(ix int, c Vec) Vec
}
