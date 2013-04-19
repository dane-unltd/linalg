package blas

func Rotg(a float64, b float64) (c float64, s float64, r float64, z float64) {
	return ops.Rotg(a, b)
}
func Rotmg(d1 float64, d2 float64, b1 float64, b2 float64) (p *RotmParams,
	rd1 float64, rd2 float64, rb1 float64) {
	return ops.Rotmg(d1, d2, b1, b2)
}
func Rotm(n int, x []float64, incX int, y []float64, incY int, p *RotmParams) {
	ops.Rotm(n, x, incX, y, incY, p)
}

func Dot(n int, x []float64, incX int, y []float64, incY int) float64 {
	return ops.Dot(n, x, incX, y, incY)
}

func Nrm2(n int, x []float64, incX int) float64 {
	return ops.Nrm2(n, x, incX)
}
func Asum(n int, x []float64, incX int) float64 {
	return ops.Asum(n, x, incX)
}
func Amax(n int, x []float64, incX int) int {
	return ops.Iamax(n, x, incX)
}
func Swap(n int, x []float64, incX int, y []float64, incY int) {
	ops.Swap(n, x, incX, y, incY)
}
func Copy(n int, x []float64, incX int, y []float64, incY int) {
	ops.Copy(n, x, incX, y, incY)
}
func Axpy(n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	ops.Axpy(n, alpha, x, incX, y, incY)
}
func Rot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	ops.Rot(n, x, incX, y, incY, c, s)
}
func Scal(n int, alpha float64, x []float64, incX int) {
	ops.Scal(n, alpha, x, incX)
}
