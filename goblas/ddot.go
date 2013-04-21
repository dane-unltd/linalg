/*
Copyright (c) 2011, Michal Derkacz
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:
1. Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.
3. The name of the author may not be used to endorse or promote products
   derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package goblas

func (Blas) Ddot(N int, X []float64, incX int, Y []float64, incY int) float64 {
	return Ddot(N, X, incX, Y, incY)
}

// Scalar product: X^T Y 
func Ddot(N int, X []float64, incX int, Y []float64, incY int) float64

func ddot(N int, X []float64, incX int, Y []float64, incY int) float64 {
	var (
		a, b, c, d float64
		xi, yi     int
	)
	for ; N >= 4; N -= 4 {
		a += X[xi] * Y[yi]
		xi += incX
		yi += incY

		b += X[xi] * Y[yi]
		xi += incX
		yi += incY

		c += X[xi] * Y[yi]
		xi += incX
		yi += incY

		d += X[xi] * Y[yi]
		xi += incX
		yi += incY
	}
	for ; N > 0; N-- {
		a += X[xi] * Y[yi]
		xi += incX
		yi += incY
	}
	return (b + c) + (d + a)
}
