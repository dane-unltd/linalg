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

TEXT ·Dasum(SB), 7, $0
	JMP	·dasum(SB)
TEXT ·Daxpy(SB), 7, $0
	JMP	·daxpy(SB)
TEXT ·Dcopy(SB), 7, $0
	JMP	·dcopy(SB)
TEXT ·Ddot(SB), 7, $0
	JMP	·ddot(SB)
TEXT ·Dnrm2(SB), 7, $0
	JMP	·dnrm2(SB)
TEXT ·Drot(SB), 7, $0
	JMP	·drot(SB)
TEXT ·Drotg(SB), 7, $0
	JMP	·drotg(SB)
TEXT ·Dscal(SB), 7, $0
	JMP	·dscal(SB)
TEXT ·Dswap(SB), 7, $0
	JMP	·dswap(SB)
TEXT ·Idamax(SB), 7, $0
	JMP	·idamax(SB)
TEXT ·Isamax(SB), 7, $0
	JMP	·isamax(SB)
TEXT ·Sasum(SB), 7, $0
	JMP	·sasum(SB)
TEXT ·Saxpy(SB), 7, $0
	JMP	·saxpy(SB)
TEXT ·Scopy(SB), 7, $0
	JMP	·scopy(SB)
TEXT ·Sdot(SB), 7, $0
	JMP	·sdot(SB)
TEXT ·Sdsdot(SB), 7, $0
	JMP	·sdsdot(SB)
TEXT ·Snrm2(SB), 7, $0
	JMP	·snrm2(SB)
TEXT ·Srot(SB), 7, $0
	JMP	·srot(SB)
TEXT ·Srotg(SB), 7, $0
	JMP	·srotg(SB)
TEXT ·Sscal(SB), 7, $0
	JMP	·sscal(SB)
TEXT ·Sswap(SB), 7, $0
	JMP	·sswap(SB)
