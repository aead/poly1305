// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// +build amd64, !gccgo, !appengine

#include "textflag.h"

#define POLY1305_ADD(msg, h0, h1, h2) \
	ADDQ 0(msg), h0; \
	ADCQ 8(msg), h1; \
	LEAQ 16(msg), msg; \
	ADCQ $1,h2
	
#define POLY1305_MUL(h0, h1, h2, r0, r1, t0, t1, R11, R12, R13, R14) \
	MOVQ r0, t0; \
	MULQ h0; \
	MOVQ t0, R11; \
	MOVQ t1, R12; \
	MOVQ r0, t0; \
	MULQ h1; \
	ADDQ t0, R12; \
	ADCQ $0, t1; \
	MOVQ r0, R13; \
	IMULQ h2, R13; \
	ADDQ t1, R13; \
    				\
	MOVQ r1, t0; \
	MULQ h0; \
	ADDQ t0, R12; \
	ADCQ $0, t1; \
	MOVQ t1, h0; \
	MOVQ r1, t0; \
	MULQ h1; \
	ADDQ h0, R13; \
	ADCQ $0, t1; \
	ADDQ t0, R13; \
	ADCQ $0, t1; \
	MOVQ r1, R14; \
	IMULQ h2, R14; \
	ADDQ t1, R14; \
					\
	MOVQ R11, h0; \
	MOVQ R12, h1; \
	MOVQ R13, h2; \
	ANDQ $3, h2; \
	MOVQ R13, R11; \
	MOVQ R14, R12; \
	ANDQ $0XFFFFFFFFFFFFFFFC, R11; \
	BYTE $0x4d; BYTE $0x0f; BYTE $0xac; BYTE $0xf5; BYTE $0x02; \ // SHRDQ $2, R14, R13
	SHRQ $2, R14; \
	ADDQ R11, h0; \
	ADCQ R12, h1; \
	ADCQ $0, h2; \
	ADDQ R13, h0; \
	ADCQ R14, h1; \
	ADCQ $0, h2

DATA poly1305Mask<>+0x00(SB)/8, $0x0FFFFFFC0FFFFFFF
DATA poly1305Mask<>+0x08(SB)/8, $0x0FFFFFFC0FFFFFFC
GLOBL poly1305Mask<>(SB), RODATA, $16

// func initialize(state *[7]uint64, key *[32]byte)
TEXT ·initialize(SB),$0-16
	MOVQ state+0(FP), DI
	MOVQ key+8(FP), SI

	MOVOU 0(SI), X0
	MOVOU 16(SI), X1
	PAND poly1305Mask<>(SB), X0
	MOVOU X0, 24(DI)
	MOVOU X1, 40(DI)
	RET

// func finalize(tag *[TagSize]byte, state *[7]uint64)
TEXT ·finalize(SB),$0-16
	MOVQ tag+0(FP),DI
	MOVQ state+8(FP),SI
	
	MOVQ 0(SI), AX
	MOVQ 8(SI), BX
	MOVQ 16(SI), CX
	MOVQ AX, R8
	MOVQ BX, R9
	MOVQ CX, R10
	SUBQ $0XFFFFFFFFFFFFFFFB, AX
	SBBQ $0XFFFFFFFFFFFFFFFF, BX
	SBBQ $3, CX
	CMOVQCS	R8, AX
	CMOVQCS	R9, BX
	CMOVQCS	R10, CX
	ADDQ 40(SI), AX
	ADCQ 48(SI), BX
	
	MOVQ AX, 0(DI)
	MOVQ BX, 8(DI)
	RET

// func core(state *[7]uint64, msg []byte)
TEXT ·core(SB),$0-32
	MOVQ state+0(FP), DI
	MOVQ msg_base+8(FP), SI
	MOVQ msg_len+16(FP), R10
	
	MOVQ 0(DI), BX		// h0
	MOVQ 8(DI), CX		// h1
	MOVQ 16(DI), R9		// h2
	MOVQ 24(DI), R15	// r0
	MOVQ 32(DI), R8 	// h1
	
	CMPQ	R10,$16
	JB BYTES_BETWEEN_0_AND_15
	
LOOP:
	POLY1305_ADD(SI, BX, CX, R9)
MULTIPLY:
	POLY1305_MUL(BX, CX, R9, R15, R8, AX, DX, R11, R12, R13, R14)
	SUBQ	$16,R10
	CMPQ	R10,$16
	JAE	LOOP
	
BYTES_BETWEEN_0_AND_15:
	TESTQ R10,R10
	JZ	DONE
	MOVQ $1, R11
	XORQ R12, R12
	XORQ R13, R13
	ADDQ R10, SI
	
FLUSH_BUFFER:
	BYTE $0x4d; BYTE $0x0f; BYTE $0xa4; BYTE $0xdc; BYTE $0x08 // SHLDQ	$8, R11, R12
	SHLQ $8, R11
	BYTE $0x4c; BYTE $0x0f; BYTE $0xb6; BYTE $0x6e; BYTE $0xff // MOVZXB -1(SI), R13
	XORQ R13, R11
	DECQ SI
	DECQ R10
	JNZ	FLUSH_BUFFER
	
	ADDQ R11, BX
	ADCQ R12, CX
	ADCQ $0, R9
	MOVQ $16, R10
	JMP	MULTIPLY	
	
DONE:
	MOVQ BX, 0(DI)
	MOVQ CX, 8(DI)
	MOVQ R9, 16(DI)
	RET
