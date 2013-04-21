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

// 6g -B: 1926 ns/op, 6a MOVUPD: 2397 ns/op, 6a MOVAPD: 474 ns/op
// MOVUPD is very inefficient on my U4100. Can it work better on other CPU?

// func Dswap(N int, X []float64, incX int, Y []float64, incY int)
TEXT ·Dswap(SB), 7, $0
	MOVL	N+0(FP), BP
	MOVQ	X_data+8(FP), SI
	MOVL	incX+24(FP), AX
	MOVQ	Y_data+32(FP), DI
	MOVL	incY+48(FP), BX

	// Check data bounaries
	MOVL	BP, CX
	DECL	CX
	MOVL	CX, DX
	IMULL	AX, CX	// CX = incX * (N - 1)
	IMULL	BX, DX	// DX = incY * (N - 1)
	CMPL	CX, X_len+16(FP)
	JGE		panic
	CMPL	DX, Y_len+40(FP)
	JGE		panic

	// Setup strides
	SALQ	$3, AX	// AX = sizeof(float64) * incX
	SALQ	$3, BX	// BX = sizeof(float64) * incY

	// Check that there are 4 or more pairs for SIMD calculations
	SUBQ	$4, BP
	JL		rest	// There are less than 4 pairs to process

	// Check if incX != 1 or incY != 1
	CMPQ	AX, $8
	JNE	with_stride
	CMPQ	BX, $8
	JNE	with_stride

	// Fully optimized loop (for incX == incY == 1)
	full_simd_loop:
		// Load four values from X
		MOVUPD	(SI), X0
		MOVUPD	16(SI), X1
		// Load four values from Y
		MOVUPD	(DI), X2
		MOVUPD	16(DI), X3
		// Save them
		MOVUPD	X0, (DI)
		MOVUPD	X1, 16(DI)
		MOVUPD	X2, (SI)
		MOVUPD	X3, 16(SI)

		// Update data pointers
		ADDQ	$32, SI
		ADDQ	$32, DI

		SUBQ	$4, BP
		JGE		full_simd_loop	// There are 4 or more pairs to process

	JMP rest

with_stride:
	// Setup long strides
	MOVQ	AX, CX
	MOVQ	BX, DX
	SALQ	$1, CX 	// CX = 16 * incX
	SALQ	$1, DX 	// DX = 16 * incY

	// Partially optimized loop
	half_simd_loop:
		// Load two values from X
		MOVSD	(SI), X0
		MOVSD	(SI)(AX*1), X1
		// Load two values from Y
		MOVSD	(DI), X2
		MOVSD	(DI)(BX*1), X3
		// Save them
		MOVSD	X0, (DI)
		MOVSD	X1, (DI)(BX*1)
		MOVSD	X2, (SI)
		MOVSD	X3, (SI)(AX*1)

		// Update data pointers using long strides
		ADDQ	CX, SI
		ADDQ	DX, DI

		// Load two values from X
		MOVSD	(SI), X0
		MOVSD	(SI)(AX*1), X1
		// Load two values from Y
		MOVSD	(DI), X2
		MOVSD	(DI)(BX*1), X3
		// Save them
		MOVSD	X0, (DI)
		MOVSD	X1, (DI)(BX*1)
		MOVSD	X2, (SI)
		MOVSD	X3, (SI)(AX*1)

		// Update data pointers using long strides
		ADDQ	CX, SI
		ADDQ	DX, DI

		SUBQ	$4, BP
		JGE		half_simd_loop	// There are 4 or more pairs to process

rest:
	// Undo last SUBQ
	ADDQ	$4,	BP

	// Check that are there any value to process
	JE	end

	loop:
		// Load values from X and Y
		MOVSD	(SI), X0
		MOVSD	(DI), X1
		// Save them
		MOVSD	X0, (DI)
		MOVSD	X1, (SI)

		// Update data pointers
		ADDQ	AX, SI
		ADDQ	BX,	DI

		DECQ	BP
		JNE	loop

end:
	RET

panic:
	CALL	runtime·panicindex(SB)
	RET
