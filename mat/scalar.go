package mat

type Scalar float64

func (d Scalar) Size() (int, int) {
	return 1, 1
}

func (d Scalar) At(i, j int) float64 {
	if i != 0 && j != 0 {
		panic("j,i have to be 0")
	}
	return float64(d)
}
