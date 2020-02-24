// Copyright 2020 The xxhash3 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !386,!amd64,!amd64p32
// +build gccgo

package cpu

func Initialize(env string) {}

func initCPU() {
	CPUID = func(op uint32) (eax, ebx, ecx, edx uint32) {
		return 0, 0, 0, 0
	}

	CPUIndex = func(op, op2 uint32) (eax, ebx, ecx, edx uint32) {
		return 0, 0, 0, 0
	}

	Xgetbv = func(index uint32) (eax, edx uint32) {
		return 0, 0
	}

	Rdtscp = func() (eax, ebx, ecx, edx uint32) {
		return 0, 0, 0, 0
	}
}
