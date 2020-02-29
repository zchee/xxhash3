// Copyright 2020 The xxhash3 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xxhash3

import (
	"testing"
	_ "unsafe"
)

//go:linkname runtime_fastrand runtime.fastrand
func runtime_fastrand() uint32

func Benchmark_avalanche(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h64 := uint64(runtime_fastrand())
		_ = avalanche(h64)
	}
}

func Benchmark_bswap32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bswap32(runtime_fastrand())
	}
}
