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

func (Blas) Daxpy(N int, alpha float64, X []float64, incX int, Y []float64, incY int) {
	Daxpy(N, alpha, X, incX, Y, incY)
}

// Compute the sum Y = \alpha X + Y for the vectors X and Y 
func Daxpy(N int, alpha float64, X []float64, incX int, Y []float64, incY int)

func daxpy(N int, alpha float64, X []float64, incX int, Y []float64, incY int) {
	var xi, yi int
	switch alpha {
	case 0:
	case 1:
		for ; N >= 2; N -= 2 {
			Y[yi] += X[xi]
			xi += incX
			yi += incY

			Y[yi] += X[xi]
			xi += incX
			yi += incY
		}
		if N != 0 {
			Y[yi] += alpha * X[xi]
		}
	case -1:
		for ; N >= 2; N -= 2 {
			Y[yi] -= X[xi]
			xi += incX
			yi += incY

			Y[yi] -= X[xi]
			xi += incX
			yi += incY
		}
		if N != 0 {
			Y[yi] -= X[xi]
		}
	default:
		for ; N >= 2; N -= 2 {
			Y[yi] += alpha * X[xi]
			xi += incX
			yi += incY

			Y[yi] += alpha * X[xi]
			xi += incX
			yi += incY
		}
		if N != 0 {
			Y[yi] += alpha * X[xi]
		}
	}
}
