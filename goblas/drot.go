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

func (Blas) Drot(N int, X []float64, incX int, Y []float64, incY int, c, s float64) {
	Drot(N, X, incX, Y, incY, c, s)
}

// Apply a Givens rotation (X', Y') = (c X + s Y, c Y - s X) to the vectors X, Y
func Drot(N int, X []float64, incX int, Y []float64, incY int, c, s float64)

func drot(N int, X []float64, incX int, Y []float64, incY int, c, s float64) {
	var xi, yi int
	for ; N > 0; N-- {
		x := X[xi]
		y := Y[yi]
		X[xi] = c*x + s*y
		Y[yi] = c*y - s*x
		xi += incX
		yi += incY
	}
}
