// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build 386 amd64

package cpu

import (
	"testing"
)

func benchIsSet(hwc uint32, value uint32) bool {
	return hwc&value != 0
}

func benchIsSet2(bitpos uint, value uint32) bool {
	return value&(1<<bitpos) != 0
}

func Benchmark_isSet(b *testing.B) {
	maxID, _, _, _ := cpuid(0, 0)
	if maxID < 1 {
		return
	}
	_, _, _, edx1 := cpuid(1, 0)

	b.ReportAllocs()
	b.ResetTimer()
	b.Run("isSet", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = benchIsSet(edx1, cpuid_SSE2)
		}
	})
}

func Benchmark_isSet2(b *testing.B) {
	maxID, _, _, _ := cpuid(0, 0)
	if maxID < 1 {
		return
	}
	_, _, _, edx1 := cpuid(1, 0)

	b.ReportAllocs()
	b.ResetTimer()
	b.Run("isSet2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = benchIsSet2(uint(edx1), cpuid_SSE2)
		}
	})
}
