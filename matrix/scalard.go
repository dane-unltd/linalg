package matrix

type ScalarD float64

func (d ScalarD) Size() (int, int) {
	return 1, 1
}

func (d ScalarD) At(i, j int) float64 {
	if i != 0 && j != 0 {
		panic("j,i have to be 0")
	}
	return float64(d)
}
