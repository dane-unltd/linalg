package blas

var Srotg func(a float32, b float32) (c float32, s float32, r float32, z float32)
var Srotmg func(d1 float32, d2 float32, b1 float32, b2 float32) (p *SrotmParams, rd1 float32, rd2 float32, rb1 float32)
var Srotm func(n int, x []float32, incX int, y []float32, incY int, p *SrotmParams)
var Drotg func(a float64, b float64) (c float64, s float64, r float64, z float64)
var Drotmg func(d1 float64, d2 float64, b1 float64, b2 float64) (p *DrotmParams, rd1 float64, rd2 float64, rb1 float64)
var Drotm func(n int, x []float64, incX int, y []float64, incY int, p *DrotmParams)

var Sdsdot func(n int, alpha float32, x []float32, incX int, y []float32, incY int) float32
var Dsdot func(n int, x []float32, incX int, y []float32, incY int) float64
var Sdot func(n int, x []float32, incX int, y []float32, incY int) float32
var Ddot func(n int, x []float64, incX int, y []float64, incY int) float64

var CdotuSub func(n int, x []complex64, incX int, y []complex64, incY int, dotu []complex64)
var CdotcSub func(n int, x []complex64, incX int, y []complex64, incY int, dotc []complex64)
var ZdotuSub func(n int, x []complex128, incX int, y []complex128, incY int, dotu []complex128)
var ZdotcSub func(n int, x []complex128, incX int, y []complex128, incY int, dotc []complex128)
var Snrm2 func(n int, x []float32, incX int) float32
var Sasum func(n int, x []float32, incX int) float32
var Dnrm2 func(n int, x []float64, incX int) float64
var Dasum func(n int, x []float64, incX int) float64
var Scnrm2 func(n int, x []complex64, incX int) float32
var Scasum func(n int, x []complex64, incX int) float32
var Dznrm2 func(n int, x []complex128, incX int) float64
var Dzasum func(n int, x []complex128, incX int) float64
var Isamax func(n int, x []float32, incX int) int
var Idamax func(n int, x []float64, incX int) int
var Icamax func(n int, x []complex64, incX int) int
var Izamax func(n int, x []complex128, incX int) int
var Sswap func(n int, x []float32, incX int, y []float32, incY int)
var Scopy func(n int, x []float32, incX int, y []float32, incY int)
var Saxpy func(n int, alpha float32, x []float32, incX int, y []float32, incY int)
var Dswap func(n int, x []float64, incX int, y []float64, incY int)
var Dcopy func(n int, x []float64, incX int, y []float64, incY int)
var Daxpy func(n int, alpha float64, x []float64, incX int, y []float64, incY int)
var Cswap func(n int, x []complex64, incX int, y []complex64, incY int)
var Ccopy func(n int, x []complex64, incX int, y []complex64, incY int)
var Caxpy func(n int, alpha complex64, x []complex64, incX int, y []complex64, incY int)
var Zswap func(n int, x []complex128, incX int, y []complex128, incY int)
var Zcopy func(n int, x []complex128, incX int, y []complex128, incY int)
var Zaxpy func(n int, alpha complex128, x []complex128, incX int, y []complex128, incY int)
var Srot func(n int, x []float32, incX int, y []float32, incY int, c float32, s float32)
var Drot func(n int, x []float64, incX int, y []float64, incY int, c float64, s float64)
var Sscal func(n int, alpha float32, x []float32, incX int)
var Dscal func(n int, alpha float64, x []float64, incX int)
var Cscal func(n int, alpha complex64, x []complex64, incX int)
var Zscal func(n int, alpha complex128, x []complex128, incX int)
var Csscal func(n int, alpha float32, x []complex64, incX int)
var Zdscal func(n int, alpha float64, x []complex128, incX int)
