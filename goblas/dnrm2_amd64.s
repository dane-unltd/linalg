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

// func Dnrm2(N int, X []float64, incX int) float64
TEXT ·Dnrm2(SB), 7, $0
	MOVL	N+0(FP), BP
	MOVQ	X+8(FP), SI	// X.data
	MOVL	incX+24(FP), AX

	// Check data bounaries
	MOVL	BP, CX
	DECL	CX
	IMULL	AX, CX	// CX = incX * (N - 1)
	CMPL	CX, X_len+16(FP)
	JGE		panic

	// Clear accumulators
	XORPD	X0, X0
	XORPD	X1, X1

	// Setup stride
	SALQ	$3, AX	// AX = sizeof(float64) * incX

	// Check that there ar 4 or more values for SIMD calculations
	SUBQ	$4, BP
	JL		rest	// There are less than 4 values to process

	// Check if incX != 1
	CMPQ	AX, $8
	JNE	with_stride

	// Fully optimized loop (for incX == incY == 1)
	full_simd_loop:
		// Multiply first two values
		MOVUPD	(SI), X2
		MULPD	X2, X2
		// Multiply second two values
		MOVUPD	16(SI), X4
		MULPD	X4, X4

		// Update data pointer
		ADDQ	$32, SI

		// Accumulate the results of multiplications
		ADDPD	X2, X0
		ADDPD	X4, X1

		SUBQ	$4, BP
		JGE		full_simd_loop	// There are 4 or more values to process

	JMP	hsum
	
with_stride:
	// Setup long stride
	MOVQ	AX, CX
	SALQ	$1, CX 	// CX = 16 * incX

	half_simd_loop:
		// Square first two values
		MOVLPD	(SI), X2
		MOVHPD	(SI)(AX*1), X2
		MULPD	X2, X2

		// Update data pointer using long stride
		ADDQ	CX, SI

		// Square second two values
		MOVLPD	(SI), X4
		MOVHPD	(SI)(AX*1), X4
		MULPD	X4, X4

		// Update data pointer using long stride
		ADDQ	CX, SI

		// Accumulate the results of multiplications
		ADDPD	X2, X0
		ADDPD	X4, X1

		SUBQ	$4, BP
		JGE		half_simd_loop	// There are 4 or more values to process

hsum:
	// Summ intermediate results from SIMD operations
	ADDPD	X0, X1
	// Horizontal sum
	MOVHLPS X1, X0
	ADDSD	X1, X0

rest:
	// Undo last SUBQ
	ADDQ	$4,	BP

	// Check that are there any value to process
	JE	end

loop:
	// Square one value
	MOVSD	(SI), X2
	MULSD	X2, X2

	// Update data pointers
	ADDQ	AX, SI

	// Accumulate the results of multiplication
	ADDSD	X2, X0

	DECQ	BP
	JNE	loop

end:
	// Return the square root of sum
	SQRTSD	X0, X0
	MOVSD	X0, r+32(FP)
	RET

panic:
	CALL	runtime·panicindex(SB)
	RET
