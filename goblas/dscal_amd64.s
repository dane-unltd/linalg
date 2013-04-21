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

// func Dscal(N int, alpha float64, X []float64, incX int)
TEXT ·Dscal(SB), 7, $0
	MOVL	N+0(FP), BP
	MOVSD	alpha+8(FP), X0
	MOVQ	X_data+16(FP), SI
	MOVL	incX+32(FP), AX

	// Check data bounaries
	MOVL	BP, CX
	DECL	CX
	IMULL	AX, CX	// CX = incX * (N - 1)
	CMPL	CX, X_len+24(FP)
	JGE		panic

	// Setup strides
	SALQ	$3, AX	// AX = sizeof(float64) * incX

	// Check that there are 4 or more values for SIMD calculations
	SUBQ	$4, BP
	JL		rest	// There are less than 4 values to process

	// Setup two alphas in X0
	MOVLHPS	X0, X0

	// Check if incX != 1
	CMPQ	AX, $8
	JNE	with_stride

	// Fully optimized loop (for incX == 1)
	full_simd_loop:
		// Load first two values and scale
		MOVUPD	(SI), X2
		MULPD	X0, X2
		// Load second two values and scale
		MOVUPD	16(SI), X4
		MULPD	X0, X4
		// Save scaled values
		MOVUPD	X2, (SI)
		MOVUPD	X4, 16(SI)

		// Update data pointers
		ADDQ	$32, SI

		SUBQ	$4, BP
		JGE		full_simd_loop	// There are 4 or more pairs to process

	JMP	rest

with_stride:
	// Setup long stride
	MOVQ	AX, CX
	SALQ	$1, CX 	// CX = 16 * incX

	// Partially optimized loop
	half_simd_loop:
		// Load first two values and scale
		MOVLPD	(SI), X2
		MOVHPD	(SI)(AX*1), X2
		MULPD	X0, X2
		// Save scaled values
		MOVLPD	X2, (SI)
		MOVHPD	X2, (SI)(AX*1)

		// Update data pointer using long stride
		ADDQ	CX, SI

		// Load second two values and scale
		MOVLPD	(SI), X4
		MOVHPD	(SI)(AX*1), X4
		MULPD	X0, X4
		// Save scaled values
		MOVLPD	X4, (SI)
		MOVHPD	X4, (SI)(AX*1)

		// Update data pointer using long stride
		ADDQ	CX, SI

		SUBQ	$4, BP
		JGE		half_simd_loop	// There are 4 or more pairs to process

rest:
	// Undo last SUBQ
	ADDQ	$4,	BP

	// Check that are there any value to process
	JE	end

	loop:
		// Load from X and scale
		MOVSD	(SI), X2
		MULSD	X0, X2
		// Save scaled value
		MOVSD	X2, (SI)

		// Update data pointers
		ADDQ	AX, SI

		DECQ	BP
		JNE	loop

end:
	RET

panic:
	CALL	runtime·panicindex(SB)
	RET
