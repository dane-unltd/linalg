package blas

func Srotg(a float32, b float32) (c float32, s float32, r float32, z float32) {
	return Ops.Srotg(a, b)
}

func Srotmg(d1 float32, d2 float32, b1 float32, b2 float32) (p *SrotmParams, rd1 float32, rd2 float32, rb1 float32) {
	return Ops.Srotmg(d1, d2, b1, b2)
}
func Srotm(n int, x []float32, incX int, y []float32, incY int, p *SrotmParams) {
	Ops.Srotm(n, x, incX, y, incY, p)
}
func Drotg(a float64, b float64) (c float64, s float64, r float64, z float64) {
	return Ops.Drotg(a, b)
}
func Drotmg(d1 float64, d2 float64, b1 float64, b2 float64) (p *DrotmParams,
	rd1 float64, rd2 float64, rb1 float64) {
	return Ops.Drotmg(d1, d2, b1, b2)
}
func Drotm(n int, x []float64, incX int, y []float64, incY int, p *DrotmParams) {
	Ops.Drotm(n, x, incX, y, incY, p)
}

func Sdsdot(n int, alpha float32, x []float32, incX int, y []float32, incY int) float32 {
	return Ops.Sdsdot(n, alpha, x, incX, y, incY)
}
func Dsdot(n int, x []float32, incX int, y []float32, incY int) float64 {
	return Ops.Dsdot(n, x, incX, y, incY)
}
func Sdot(n int, x []float32, incX int, y []float32, incY int) float32 {
	return Ops.Sdot(n, x, incX, y, incY)
}
func Ddot(n int, x []float64, incX int, y []float64, incY int) float64 {
	return Ops.Ddot(n, x, incX, y, incY)
}

func Cdotu(n int, x []complex64, incX int, y []complex64, incY int) complex64 {
	return Ops.Cdotu(n, x, incX, y, incY)
}
func Cdotc(n int, x []complex64, incX int, y []complex64, incY int) complex64 {
	return Ops.Cdotc(n, x, incX, y, incY)
}
func Zdotu(n int, x []complex128, incX int, y []complex128, incY int) complex128 {
	return Ops.Zdotu(n, x, incX, y, incY)
}
func Zdotc(n int, x []complex128, incX int, y []complex128, incY int) complex128 {
	return Ops.Zdotc(n, x, incX, y, incY)
}
func Snrm2(n int, x []float32, incX int) float32 {
	return Ops.Snrm2(n, x, incX)
}
func Sasum(n int, x []float32, incX int) float32 {
	return Ops.Sasum(n, x, incX)
}
func Dnrm2(n int, x []float64, incX int) float64 {
	return Ops.Dnrm2(n, x, incX)
}
func Dasum(n int, x []float64, incX int) float64 {
	return Ops.Dasum(n, x, incX)
}
func Scnrm2(n int, x []complex64, incX int) float32 {
	return Ops.Scnrm2(n, x, incX)
}
func Scasum(n int, x []complex64, incX int) float32 {
	return Ops.Scasum(n, x, incX)
}
func Dznrm2(n int, x []complex128, incX int) float64 {
	return Ops.Dznrm2(n, x, incX)
}
func Dzasum(n int, x []complex128, incX int) float64 {
	return Ops.Dzasum(n, x, incX)
}
func Isamax(n int, x []float32, incX int) int {
	return Ops.Isamax(n, x, incX)
}
func Idamax(n int, x []float64, incX int) int {
	return Ops.Idamax(n, x, incX)
}
func Icamax(n int, x []complex64, incX int) int {
	return Ops.Icamax(n, x, incX)
}
func Izamax(n int, x []complex128, incX int) int {
	return Ops.Izamax(n, x, incX)
}
func Sswap(n int, x []float32, incX int, y []float32, incY int) {
	Ops.Sswap(n, x, incX, y, incY)
}
func Scopy(n int, x []float32, incX int, y []float32, incY int) {
	Ops.Scopy(n, x, incX, y, incY)
}
func Saxpy(n int, alpha float32, x []float32, incX int, y []float32, incY int) {
	Ops.Saxpy(n, alpha, x, incX, y, incY)
}
func Dswap(n int, x []float64, incX int, y []float64, incY int) {
	Ops.Dswap(n, x, incX, y, incY)
}
func Dcopy(n int, x []float64, incX int, y []float64, incY int) {
	Ops.Dcopy(n, x, incX, y, incY)
}
func Daxpy(n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	Ops.Daxpy(n, alpha, x, incX, y, incY)
}
func Cswap(n int, x []complex64, incX int, y []complex64, incY int) {
	Ops.Cswap(n, x, incX, y, incY)
}
func Ccopy(n int, x []complex64, incX int, y []complex64, incY int) {
	Ops.Ccopy(n, x, incX, y, incY)
}
func Caxpy(n int, alpha complex64, x []complex64, incX int, y []complex64, incY int) {
	Ops.Caxpy(n, alpha, x, incX, y, incY)
}
func Zswap(n int, x []complex128, incX int, y []complex128, incY int) {
	Ops.Zswap(n, x, incX, y, incY)
}
func Zcopy(n int, x []complex128, incX int, y []complex128, incY int) {
	Ops.Zcopy(n, x, incX, y, incY)
}
func Zaxpy(n int, alpha complex128, x []complex128, incX int, y []complex128, incY int) {
	Ops.Zaxpy(n, alpha, x, incX, y, incY)
}
func Srot(n int, x []float32, incX int, y []float32, incY int, c float32, s float32) {
	Ops.Srot(n, x, incX, y, incY, c, s)
}
func Drot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	Ops.Drot(n, x, incX, y, incY, c, s)
}
func Sscal(n int, alpha float32, x []float32, incX int) {
	Ops.Sscal(n, alpha, x, incX)
}
func Dscal(n int, alpha float64, x []float64, incX int) {
	Ops.Dscal(n, alpha, x, incX)
}
func Cscal(n int, alpha complex64, x []complex64, incX int) {
	Ops.Cscal(n, alpha, x, incX)
}
func Zscal(n int, alpha complex128, x []complex128, incX int) {
	Ops.Zscal(n, alpha, x, incX)
}
func Csscal(n int, alpha float32, x []complex64, incX int) {
	Ops.Csscal(n, alpha, x, incX)
}
func Zdscal(n int, alpha float64, x []complex128, incX int) {
	Ops.Zdscal(n, alpha, x, incX)
}
