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

import "math"

func (Blas) Drotg(a, b float64) (c, s, r, z float64) {
	return Drotg(a, b)
}

// Compute a Givens rotation (c,s) which zeroes the vector (a,b)
func Drotg(a, b float64) (c, s, r, z float64)

func drotg(a, b float64) (c, s, r, z float64) {
	abs_a := a
	if a < 0 {
		abs_a = -a
	}
	abs_b := b
	if b < 0 {
		abs_b = -b
	}
	roe := b
	if abs_a > abs_b {
		roe = a
	}
	scale := abs_a + abs_b
	if scale == 0 {
		c = 1
	} else {
		sa := a / scale
		sb := b / scale
		r = scale * math.Sqrt(sa*sa+sb*sb)
		if roe < 0 {
			r = -r
		}
		c = a / r
		s = b / r
		z = 1
		if abs_a > abs_b {
			z = s
		} else {
			if c != 0 {
				z /= c
			}
		}
	}
	return
}
