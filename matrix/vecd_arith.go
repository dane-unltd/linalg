package matrix

func (v VecD) Nrm2Sq() float64 {
	res := 0.0
	for _, val := range v {
		res += val * val
	}
	return res
}
