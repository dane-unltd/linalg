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

// func Idamax(N int, X []float64, incX int) int
TEXT ·Idamax(SB), 7, $0
	MOVL	N+0(FP), BP
	MOVQ	X+8(FP), SI	// X.data
	MOVL	incX+24(FP), AX

	// Check data bounaries
	MOVL	BP, CX
	DECL	CX
	IMULL	AX, CX	// CX = incX * (N - 1)
	CMPL	CX, X_len+16(FP)
	JGE		panic

	// Max value
	XORPD	X0, X0
	// Index
	XORQ	DI, DI
	// Max index
	XORQ	BX, BX

	// Mask for sign bit clear
	PCMPEQL	X4, X4 
	PSRLQ	$1, X4

	// Setup stride
	SALQ	$3, AX	// AX = sizeof(float64) * incX

loop:
	CMPQ	BP, DI
	JE	end

	// Load value
	MOVSD	(SI), X1
	// Clear sign of loaded value
	ANDPD	X4, X1

	// Is loaded value less or equal to max value?
	UCOMISD	X0,	X1
	JBE	less_or_equal

	// Save max index and value
	MOVQ	DI, BX
	MOVSD	X1, X0

less_or_equal:
	// Update data pointers
	ADDQ	AX, SI
	INCQ	DI
	JMP	loop

end:
	// Return the max index
	MOVL	BX, r+32(FP)
	RET

panic:
	CALL	runtime·panicindex(SB)
	RET
