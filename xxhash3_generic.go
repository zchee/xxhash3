// Copyright 2020 The xxhash3 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !amd64

package xxhash3

var (
	useSSE2 = false
	useAVX2 = false
	useNEON = false
	useVMX  = false
)

func init() {
	Align = 8
}
