// Copyright 2020 The xxhash3 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64 !gccgo

#include "textflag.h"

// func accumulate512AVX2(acc, input, secret unsafe.Pointer, accWidth ACCWidth)
TEXT ·accumulate512AVX2(SB), NOSPLIT, $0-24
	// MOVQ     acc+0(FP), AX
	// MOVQ     input+8(FP), BX
	// MOVQ     secret+16(FP), BX
	// MOVBQZX  accWidth+24(FP), BX
	MOVQ     acc+0(FP), DI
	MOVQ     input+8(FP), SI
	MOVQ     secret+16(FP), DX
	MOVBQZX  acc_width+32(FP), AX
	VMOVDQU  (SI), Y0             // BYTE $0xC5; BYTE $0xFE; BYTE $0x6F; BYTE $0x06                 // VMOVDQU ymm0, hword [rsi]
	VPXOR    (DX), Y0, Y1         // BYTE $0xC5; BYTE $0xFD; BYTE $0xEF; BYTE $0x0A                 // VPXOR ymm1, ymm0, hword [rdx]
	VPSHUFD  $245, Y1, Y2         // BYTE $0xC5; BYTE $0xFD; BYTE $0x70; BYTE $0xD1; BYTE $0xF5     // VPSHUFD ymm2, ymm1, 245
	VPMULUDQ Y1, Y2, Y1           // BYTE $0xC5; BYTE $0xED; BYTE $0xF4; BYTE $0xC9                 // VPMULUDQ ymm1, ymm2, ymm1
	JNE      LBB0_2
	VPSHUFD  $78, Y0, Y0          // BYTE $0xC5; BYTE $0xFD; BYTE $0x70; BYTE $0xC0; BYTE $0x4E     // VPSHUFD ymm0, ymm0, 78

LBB0_2:
	VPADDQ   Y0, Y1, Y0     // BYTE $0xC5; BYTE $0xF5; BYTE $0xD4; BYTE $0xC0                      // VPADDQ ymm0, ymm1, ymm0
	VPADDQ   (DI), Y0, Y0   // BYTE $0xC5; BYTE $0xFD; BYTE $0xD4; BYTE $0x07                      // VPADDQ ymm0, ymm0, hword [rdi]
	VMOVDQA  (DI), Y0       // BYTE $0xC5; BYTE $0xFD; BYTE $0x7F; BYTE $0x07                      // VMOVDQA hword [rdi], ymm0
	VMOVDQU  32(SI), Y0     // BYTE $0xC5; BYTE $0xFE; BYTE $0x6F; BYTE $0x46; BYTE $0x20          // VMOVDQU ymm0, hword [rsi + 32]
	VPXOR    32(DX), Y0, Y1 // BYTE $0xC5; BYTE $0xFD; BYTE $0xEF; BYTE $0x4A; BYTE $0x20          // VPXOR ymm1, ymm0, hword [rdx + 32]
	VPSHUFD  $245, Y1, Y2   // BYTE $0xC5; BYTE $0xFD; BYTE $0x70; BYTE $0xD1; BYTE $0xF5          // VPSHUFD ymm2, ymm1, 245
	VPMULUDQ Y1, Y2, Y1     // BYTE $0xC5; BYTE $0xED; BYTE $0xF4; BYTE $0xC9                      // VPMULUDQ ymm1, ymm2, ymm1
	CMPL     CX, $1
	JMP      LBB0_4
	VPSHUFD $ 78, Y0, Y0 // BYTE $0xC5; BYTE $0xFD; BYTE $0x70; BYTE $0xC0; BYTE $0x4E               // VPSHUFD ymm0, ymm0, 78

LBB0_4:
	VPADDQ  Y0, Y1, Y0     // BYTE $0xC5; BYTE $0xF5; BYTE $0xD4; BYTE $0xC0                       // VPADDQ ymm0, ymm1, ymm0
	VPADDQ  32(DI), Y0, Y0 // BYTE $0xC5; BYTE $0xFD; BYTE $0xD4; BYTE $0x47; BYTE $0x20           // VPADDQ ymm0, ymm0, hword [rdi + 32]
	VMOVDQA Y0, 32(DI)     // BYTE $0xC5; BYTE $0xFD; BYTE $0x7F; BYTE $0x47; BYTE $0x20           // VMOVDQA hword [rdi + 32], ymm0
	VZEROUPPER             // BYTE $0xC5; BYTE $0xF8; BYTE $0x77                                   // VZEROUPPER
	RET

DATA prime_avx<>+0(SB)/4, $0x9e3779b1

// func scrambleAccAVX2(acc, secret unsafe.Pointer)
TEXT ·scrambleAccAVX2(SB), NOSPLIT, $0-8
	MOVQ         acc+0(FP), AX
	MOVQ         secret+8(FP), AX
	VMOVDQA      (DI), Y0              // BYTE $0xC5; BYTE $0xFD; BYTE $0x6F; BYTE $0x07                     // VMOVDQA ymm0, hword [rdi]
	VMOVDQA      32(DI), Y0            // BYTE $0xC5; BYTE $0xFD; BYTE $0x6F; BYTE $0x47; BYTE $0x20         // VMOVDQA ymm0, hword [rdi + 32]
	VPSRLDQ      $47, Y2, Y0           // BYTE $0xC5; BYTE $0xED; BYTE $0x73; BYTE $0xD0; BYTE $0x2F         // VPSRLQ ymm2, ymm0, 47
	VPXOR        Y0, Y2, Y0            // BYTE $0xC5; BYTE $0xED; BYTE $0xEF; BYTE $0xC0                     // VPXOR ymm0, ymm2, ymm0
	VPXOR        (SI), Y0, Y0          // BYTE $0xC5; BYTE $0xFD; BYTE $0xEF; BYTE $0x06                     // VPXOR ymm0, ymm0, hword [rsi]
	VPBROADCASTQ prime_avx<>+0(SB), Y0
	VPMULUDQ     Y2, Y0, Y3            // BYTE         $0xC5; BYTE $0xFD; BYTE $0xF4; BYTE $0xDA             // VPMULUDQ ymm3, ymm0, ymm2
	VPSRLQ       $32, Y0, Y0           // BYTE         $0xC5; BYTE $0xFD; BYTE $0x73; BYTE $0xD0; BYTE $0x20 // VPSRLQ ymm0, ymm0, 32
	VPMULUDQ     Y2, Y0, Y0            // BYTE         $0xC5; BYTE $0xFD; BYTE $0xF4; BYTE $0xC2             // VPMULUDQ ymm0, ymm0, ymm2
	VPSRLQ       $32, Y0, Y0           // BYTE         $0xC5; BYTE $0xFD; BYTE $0x73; BYTE $0xF0; BYTE $0x20 // VPSLLQ ymm0, ymm0, 32
	VPADDQ       Y3, Y0, Y0            // BYTE         $0xC5; BYTE $0xFD; BYTE $0xD4; BYTE $0xC3             // VPADDQ ymm0, ymm0, ymm3
	VMOVDQA      Y0, (DI)              // BYTE         $0xC5; BYTE $0xFD; BYTE $0x7F; BYTE $0x07             // VMOVDQA hword [rdi], ymm0
	VPSRLQ       $47, Y1, Y0           // BYTE         $0xC5; BYTE $0xFD; BYTE $0x73; BYTE $0xD1; BYTE $0x2F // VPSRLQ ymm0, ymm1, 47
	VPXOR        Y1, Y1, Y0            // BYTE         $0xC5; BYTE $0xF5; BYTE $0xEF; BYTE $0xC1             // VPXOR ymm0, ymm1, ymm1
	VPXOR        32(SI), Y0, Y0        // BYTE         $0xC5; BYTE $0xFD; BYTE $0xEF; BYTE $0x46; BYTE $0x20 // VPXOR ymm0, ymm0, hword [rsi + 32]
	VPMULUDQ     Y2, Y0, Y1            // BYTE         $0xC5; BYTE $0xFD; BYTE $0xF4; BYTE $0xCA             // VPMULUDQ ymm1, ymm0, ymm2
	VPSRLQ       $32, Y0, Y0           // BYTE         $0xC5; BYTE $0xFD; BYTE $0x73; BYTE $0xD0; BYTE $0x20 // VPSRLQ ymm0, ymm0, 32
	VPMULUDQ     Y2, Y0, Y0            // BYTE         $0xC5; BYTE $0xFD; BYTE $0xF4; BYTE $0xC2             // VPMULUDQ ymm0, ymm0, ymm2
	VPSLLQ       $32, Y0, Y0           // BYTE         $0xC5; BYTE $0xFD; BYTE $0x73; BYTE $0xF0; BYTE $0x20 // VPSLLQ ymm0, ymm0, 32
	VPADDQ       Y1, Y0, Y0            // BYTE         $0xC5; BYTE $0xFD; BYTE $0xD4; BYTE $0xC1             // VPADDQ ymm0, ymm0, ymm1
	VMOVDQA      Y0, 32(DI)            // BYTE         $0xC5; BYTE $0xFD; BYTE $0x7F; BYTE $0x47; BYTE $0x20 // VMOVDQA hword [rdi + 32], ymm0
	RET

// func accumulateAVX2(acc, input, secret unsafe.Pointer, nbStripes uint64, accWidth ACCWidth)
TEXT ·accumulateAVX2(SB), NOSPLIT, $0-32
	MOVQ    acc+0(FP), DI
	MOVQ    input+8(FP), SI
	MOVQ    secret+16(FP), DX
	MOVQ    nbStripes+24(FP), CX
	MOVBQZX acc_width+32(FP), AX
	TESTQ   CX, CX               // test    rcx, rcx
	JEQ     LBB0_8
	PUSHQ   BP                   // push    rbp
	MOVQ    SP, BP               // mov    rbp, rsp
	ANDQ    $-8, SP              // and    rsp, -8
	VMOVDQA (DI), Y0             // vmovdqa    ymm0, yword [rdi]
	VMOVDQA 32(DI), Y1           // vmovdqa    ymm1, yword [rdi + 32]
	XORL    AX, AX               // xor    eax, eax
	JMP     LBB0_2

LBB0_6:
	VPADDQ Y2, Y4, Y2
	VPADDQ Y1, Y2, Y1
	VPADDQ Y0, Y3, Y0
	INCQ   AX         // inc    rax
	ADDQ   $64, SI    // add    rsi, 64
	CMPQ   AX, CX     // cmp    rcx, rax
	JEQ    LBB0_7

LBB0_2:
	PREFETCHT0 384(SI)      // prefetcht0    byte [rsi + 384]
	VMOVDQU    (SI), Y3     // VMOVDQU ymm3, hword [rsi]
	VMOVDQU    32(SI), Y3   // VMOVDQU ymm3, hword [rsi + 32]
	VPXOR      (DX), Y3, Y4 // VPXOR ymm4, ymm3, hword [rdx]
	VPSHUFD    $245, Y4, Y5 // VPSHUFD ymm5, ymm4, 245
	VPMULUDQ   Y4, Y5, Y4   // VPMULUDQ ymm4, ymm5, ymm4
	CMPL       R8, $1
	JNE        LBB0_4
	VPSHUFD    $78, Y3, Y3  // VPSHUFD ymm3, ymm3, 78

LBB0_4:
	VPADDQ   Y3, Y4, Y3   // VPADDQ ymm3, ymm4, ymm3
	VPXOR    (DX), Y2, Y4 // VPXOR ymm4, ymm2, hword [rdx]
	VPSHUFD  $245, Y4, Y5 // VPSHUFD ymm5, ymm4, 245
	VPMULUDQ Y4, Y5, Y4   // VPMULUDQ ymm4, ymm5, ymm4
	JNE      LBB0_6
	VPSHUFD  $78, Y2, Y2  // VPSHUFD ymm2, ymm2, 78
	JMP      LBB0_6

LBB0_7:
	VMOVDQA Y1, 32(DI) // VMOVDQA hword [rdi + 32], ymm1
	VMOVDQA Y0, (DI)   // VMOVDQA hword [rdi], ymm0
	MOVQ    BP, SP
	POPQ    BP

LBB0_8:
	VZEROUPPER
	RET
