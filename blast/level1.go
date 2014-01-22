package blast

func checkVV(x, y Vector) int {
	if x.N != y.N {
		panic(ErrShape)
	}
	return x.N
}

func Dot(x Vector, y Vector) float64 {
	must(x.Check())
	must(y.Check())

	s := 0.0
	n := checkVV(x, y)
	for i, j := 0, 0; i < n*x.Inc; i, j = i+x.Inc, j+y.Inc {
		s += x.Data[i] * y.Data[j]
	}
	return s
}

func Axpby(alpha float64, x Vector, beta float64, y Vector) {
	if alpha == 0 {
		if beta == 1 {
			return
		}
		must(y.Check())
		for j := 0; j < y.N*y.Inc; j = j + y.Inc {
			y.Data[j] = beta * y.Data[j]
		}
		return
	}
	must(x.Check())
	must(y.Check())
	n := checkVV(x, y)
	for i, j := 0, 0; i < n*x.Inc; i, j = i+x.Inc, j+y.Inc {
		y.Data[j] = beta*y.Data[j] + alpha*x.Data[i]
	}
}

func Waxpby(alpha float64, x Vector, beta float64, y, w Vector) {
	must(x.Check())
	must(y.Check())
	must(w.Check())

	n := checkVV(w, y)
	checkVV(w, x)

	for i, j, k := 0, 0, 0; i < n*x.Inc; i, j, k = i+x.Inc, j+y.Inc, k+w.Inc {
		w.Data[k] = beta*y.Data[j] + alpha*x.Data[i]
	}
}

func Copy(x, y Vector) {
	must(x.Check())
	must(y.Check())

	n := checkVV(x, y)
	for i, j := 0, 0; i < n*x.Inc; i, j = i+x.Inc, j+y.Inc {
		y.Data[j] = x.Data[i]
	}
}
