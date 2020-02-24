// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build 386 !gccgo,amd64 !gccgo,amd64p32 !gccgo

package cpu

func init() {
	detectFeatures() // intel
}

// Initialize examines the processor and sets the relevant variables above.
// This is called by the runtime package early in program initialization,
// before normal init functions are run. env is set by runtime if the OS supports
// cpu feature options in GODEBUG.
func Initialize(env string) {
	doinit()
	processOptions(env) // options
}

func asmCPUID(op uint32) (eax, ebx, ecx, edx uint32)
func asmCPUIndex(op, op2 uint32) (eax, ebx, ecx, edx uint32)
func asmXgetbv(index uint32) (eax, edx uint32)
func asmRdtscp() (eax, ebx, ecx, edx uint32)

func initCPU() {
	CPUID = asmCPUID
	CPUIndex = asmCPUIndex
	Xgetbv = asmXgetbv
	Rdtscp = asmRdtscp
}
