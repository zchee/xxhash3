// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64 amd64p32
// +build !gccgo

#include "textflag.h"

// go
// func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)
TEXT ·cpuid(SB), NOSPLIT, $0-24
	MOVL eaxArg+0(FP), AX
	MOVL ecxArg+4(FP), CX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

// go
// func xgetbv() (eax, edx uint32)
TEXT ·xgetbv(SB), NOSPLIT, $0-8
	MOVL $0, CX
	XGETBV
	MOVL AX, eax+0(FP)
	MOVL DX, edx+4(FP)
	RET

// intel-go
// func cpuid_low(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)
TEXT ·cpuid_low(SB), NOSPLIT, $0-24
	MOVL eaxArg+0(FP), AX
	MOVL ecxArg+4(FP), CX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

// intel-go
// func xgetbv_low(index uint32) (eax, edx uint32)
TEXT ·xgetbv_low(SB), NOSPLIT, $0-16
	MOVL index+0(FP), CX
	BYTE $0x0f; BYTE $0x01; BYTE $0xd0 // XGETBV
	MOVL AX, eax+8(FP)
	MOVL DX, edx+12(FP)
	RET

// gvisor
// func hostID(raxArg, rcxArg uint32) (eax, ebx, ecx, edx uint32)
TEXT ·hostID(SB),$0-48
	MOVL raxArg+0(FP), AX
	MOVL rcxArg+4(FP), CX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

// klauspost
// func asmCPUID(op uint32) (eax, ebx, ecx, edx uint32)
TEXT ·asmCPUID(SB), 7, $0
	XORQ CX, CX
	MOVL op+0(FP), AX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

// klauspost
// func CPUIndex(op, op2 uint32) (eax, ebx, ecx, edx uint32)
TEXT ·asmCPUIndex(SB), 7, $0
	MOVL op+0(FP), AX
	MOVL op2+4(FP), CX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

// klauspost
// func asmXgetbv(index uint32) (eax, edx uint32)
TEXT ·asmXgetbv(SB), 7, $0
	MOVL index+0(FP), CX
	BYTE $0x0f; BYTE $0x01; BYTE $0xd0 // XGETBV
	MOVL AX, eax+8(FP)
	MOVL DX, edx+12(FP)
	RET

// klauspost
// func asmRdtscp() (eax, ebx, ecx, edx uint32)
TEXT ·asmRdtscp(SB), 7, $0
	BYTE $0x0F; BYTE $0x01; BYTE $0xF9 // RDTSCP
	MOVL AX, eax+0(FP)
	MOVL BX, ebx+4(FP)
	MOVL CX, ecx+8(FP)
	MOVL DX, edx+12(FP)
	RET
