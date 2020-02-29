// Copyright 2020 The xxhash3 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64 !gccgo

package xxhash3

const (
	versionMajor   = "0" // #define XXH_VERSION_MAJOR    0
	versionMinor   = "7" // #define XXH_VERSION_MINOR    7
	versionRelease = "3" // #define XXH_VERSION_RELEASE  3
)

// #define XXH_VERSION_NUMBER  (XXH_VERSION_MAJOR *100*100 + XXH_VERSION_MINOR *100 + XXH_VERSION_RELEASE)
var versionNumber = "v" + versionMajor + "." + versionMinor + "." + versionRelease

// Version returns the xxhash3 current version number.
func Version() string {
	return versionNumber
}
