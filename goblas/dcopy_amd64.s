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

// func Dcopy(N int, X []float64, incX int, Y []float64, incY int)
TEXT ·Dcopy(SB), 7, $0
	MOVL	N+0(FP), CX
	MOVQ	X_data+8(FP), SI
	MOVL	incX+24(FP), AX
	MOVQ	Y_data+32(FP), DI
	MOVL	incY+48(FP), BX

	// Check data bounaries
	MOVL	CX, BP
	DECL	BP
	MOVL	BP, DX
	IMULL	AX, BP	// BP = incX * (N - 1)
	IMULL	BX, DX	// DX = incY * (N - 1)
	CMPL	BP, X_len+16(FP)
	JGE		panic
	CMPL	DX, Y_len+40(FP)
	JGE		panic

	// Check if incX != 1 or incY != 1
	CMPQ	AX, $1
	JNE	with_stride
	CMPQ	BX, $1
	JNE	with_stride

	// Optimized copy for incX == incY == 1
	REP; MOVSQ
	RET

with_stride:
	// Setup strides
	SALQ	$3, AX	// AX = sizeof(float64) * incX
	SALQ	$3, BX	// BX = sizeof(float64) * incY

	CMPQ	CX, $0
	JE	end

	loop:
		MOVQ	(SI), DX
		MOVQ	DX, (DI)
		ADDQ	AX, SI
		ADDQ	BX, DI
		DECQ	CX
		JNE	loop

end:
	RET

panic:
	CALL	runtime·panicindex(SB)
	RET
