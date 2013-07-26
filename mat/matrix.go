package mat

type Matrix interface {
	At(i, j int) float64
	Size() (int, int)
}
