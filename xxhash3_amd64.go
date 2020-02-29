// Copyright 2020 The xxhash3 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64 !gccgo

package xxhash3

import (
	"math/bits"
	"unsafe"

	"github.com/zchee/xxhash3/internal/cpu"
)

var (
	useSSE2 = cpu.X86.HasSSE2
	useAVX2 = cpu.X86.HasAVX2
	useNEON = cpu.ARM.HasNEON
	useVMX  = false
)

// vector Vectorization detection.
type vector uint8

const (
	vectorScalar vector = 1 + iota /* Portable scalar version */
	vectorSSE2                     /* SSE2 for Pentium 4 and all x86_64 */
	vectorAVX2                     /* AVX2 for Haswell and Bulldozer */
	vectorNEON                     /* NEON for most ARMv7-A and all AArch64 */
	vectorVSX                      /* VSX and ZVector for POWER8/z13 */
)

var (
	Vector vector
	Align  int
)

func init() {
	switch {
	case useSSE2:
		Vector = vectorSSE2
	case useAVX2:
		Vector = vectorAVX2
	case useNEON:
		Vector = vectorNEON
	default:
		Vector = vectorScalar
	}
	Align = 8 * int(Vector)
}

const (
	SecretSizeMin     = 136
	SecretDefaultSize = 192 // minimum SecretSizeMin
)

var kSecret = unsafe.Pointer(&[SecretDefaultSize]uint8{
	0xb8, 0xfe, 0x6c, 0x39, 0x23, 0xa4, 0x4b, 0xbe, 0x7c, 0x01, 0x81, 0x2c, 0xf7, 0x21, 0xad, 0x1c,
	0xde, 0xd4, 0x6d, 0xe9, 0x83, 0x90, 0x97, 0xdb, 0x72, 0x40, 0xa4, 0xa4, 0xb7, 0xb3, 0x67, 0x1f,
	0xcb, 0x79, 0xe6, 0x4e, 0xcc, 0xc0, 0xe5, 0x78, 0x82, 0x5a, 0xd0, 0x7d, 0xcc, 0xff, 0x72, 0x21,
	0xb8, 0x08, 0x46, 0x74, 0xf7, 0x43, 0x24, 0x8e, 0xe0, 0x35, 0x90, 0xe6, 0x81, 0x3a, 0x26, 0x4c,
	0x3c, 0x28, 0x52, 0xbb, 0x91, 0xc3, 0x00, 0xcb, 0x88, 0xd0, 0x65, 0x8b, 0x1b, 0x53, 0x2e, 0xa3,
	0x71, 0x64, 0x48, 0x97, 0xa2, 0x0d, 0xf9, 0x4e, 0x38, 0x19, 0xef, 0x46, 0xa9, 0xde, 0xac, 0xd8,
	0xa8, 0xfa, 0x76, 0x3f, 0xe3, 0x9c, 0x34, 0x3f, 0xf9, 0xdc, 0xbb, 0xc7, 0xc7, 0x0b, 0x4f, 0x1d,
	0x8a, 0x51, 0xe0, 0x4b, 0xcd, 0xb4, 0x59, 0x31, 0xc8, 0x9f, 0x7e, 0xc9, 0xd9, 0x78, 0x73, 0x64,
	0xea, 0xc5, 0xac, 0x83, 0x34, 0xd3, 0xeb, 0xc3, 0xc5, 0x81, 0xa0, 0xff, 0xfa, 0x13, 0x63, 0xeb,
	0x17, 0x0d, 0xdd, 0x51, 0xb7, 0xf0, 0xda, 0x49, 0xd3, 0x16, 0x55, 0x26, 0x29, 0xd4, 0x68, 0x9e,
	0x2b, 0x16, 0xbe, 0x58, 0x7d, 0x47, 0xa1, 0xfc, 0x8f, 0xf8, 0xb8, 0xd1, 0x7a, 0xd0, 0x31, 0xce,
	0x45, 0xcb, 0x3a, 0x8f, 0x95, 0x16, 0x04, 0x28, 0xaf, 0xd7, 0xfb, 0xca, 0xbb, 0x4b, 0x40, 0x7e,
})

type uint128 struct {
	lo uint64
	hi uint64
}

// 32-bit hash functions.
const (
	Prime32_1 = uint32(2654435761) // 0b10011110001101110111100110110001
	Prime32_2 = uint32(2246822519) // 0b10000101111010111100101001110111
	Prime32_3 = uint32(3266489917) // 0b11000010101100101010111000111101
	Prime32_4 = uint32(668265263)  // 0b00100111110101001110101100101111
	Prime32_5 = uint32(374761393)  // 0b00010110010101100110011110110001
)

// 64-bit hash functions.
const (
	Prime64_1 = uint64(11400714785074694791) // 0b1001111000110111011110011011000110000101111010111100101010000111
	Prime64_2 = uint64(14029467366897019727) // 0b1100001010110010101011100011110100100111110101001110101101001111
	Prime64_3 = uint64(1609587929392839161)  // 0b0001011001010110011001111011000110011110001101110111100111111001
	Prime64_4 = uint64(9650029242287828579)  // 0b1000010111101011110010100111011111000010101100101010111001100011
	Prime64_5 = uint64(2870177450012600261)  // 0b0010011111010100111010110010111100010110010101100110011111000101
)

func mult64to128(lhs, rhs uint64) uint128 {
	hi, lo := bits.Mul64(lhs, rhs)
	u128 := uint128{
		lo: lo,
		hi: hi,
	}
	return u128
}

func mul128Fold64(lhs, rhs uint64) uint64 {
	product := mult64to128(lhs, rhs)
	return product.lo ^ product.hi
}

func avalanche(h64 uint64) uint64 {
	h64 ^= h64 >> 37
	h64 *= Prime64_3
	h64 ^= h64 >> 32
	return h64
}

func readLE32(p unsafe.Pointer) uint32 {
	return cpu.HostByteOrder().Uint32((*[4]byte)(p)[:])
}

func readLE64(p unsafe.Pointer) uint64 {
	return cpu.HostByteOrder().Uint64((*[8]byte)(p)[:])
}

// Short keys
func len1to364b(input unsafe.Pointer, length uintptr, secret unsafe.Pointer, seed uint64) uint64 {
	c1 := *(*uint8)(input)
	c2 := *(*uint8)(unsafe.Pointer(uintptr(input) + uintptr(length>>1)))
	c3 := *(*uint8)(unsafe.Pointer(uintptr(input) + uintptr(length-1)))
	combined := uint32(uint32(c1) | (uint32(c2) << 8) | (uint32(c3) << 16) | (uint32(length) << 24))
	keyed := uint64(combined ^ (readLE32(unsafe.Pointer(secret)) + uint32(seed)))
	mixed := uint64(uint64(keyed) * Prime64_1)

	return avalanche(mixed)
}

func len4to864b(input unsafe.Pointer, length uintptr, secret unsafe.Pointer, seed uint64) uint64 {
	seed ^= uint64(bswap32(uint32(seed))) << 32

	input1 := uint32(readLE32(input))
	input2 := uint32(readLE32(unsafe.Pointer(uintptr(input) + uintptr(length-4))))
	key1 := uint32(bswap32(input1) ^ uint32(uint32((seed>>32))+readLE32(secret)))
	key2 := uint32(input2 ^ (readLE32(unsafe.Pointer(uintptr(secret)-4)) - uint32(seed)))
	hi, lo := bits.Mul64(uint64(key1), uint64(key2))
	mix64 := uint64(lo + hi + uint64(input1)<<32 + uint64(bits.RotateLeft32(input2, 23))<<32 + uint64(length))

	return avalanche(mix64 ^ (mix64 >> 59))
}

func len9to1664b(input unsafe.Pointer, length uintptr, secret unsafe.Pointer, seed uint64) uint64 {
	inputLo := readLE64(input) ^ readLE64(secret)
	inputHi := readLE64(unsafe.Pointer(uintptr(input)+uintptr(length-8))) ^ readLE64(unsafe.Pointer(uintptr(secret)+8-uintptr(seed)))
	acc := uint64(length) + (inputLo + inputHi) + mul128Fold64(inputLo, inputHi)

	return avalanche(acc)
}

func len0to1664b(input unsafe.Pointer, length uintptr, secret unsafe.Pointer, seed uint64) uint64 {
	if length > 8 {
		return len9to1664b(input, length, secret, seed)
	}
	if length >= 4 {
		return len4to864b(input, length, secret, seed)
	}
	if length > 0 {
		return len1to364b(input, length, secret, seed)
	}

	return avalanche((Prime64_1 + seed) ^ readLE64(secret))
}

// Long Keys
const (
	StripeLen         = 64
	SecretConsumeRate = 8
	ACCNB             = StripeLen / unsafe.Sizeof(new(uint64))
)

type ACCWidth uint8

const (
	acc64bits ACCWidth = iota
	acc128bits
)

//go:noescape
func accumulate512AVX2(acc, input, secret unsafe.Pointer, accWidth ACCWidth)

func accumulate512(acc, input, secret unsafe.Pointer, accWidth ACCWidth) {
	switch {
	case useAVX2:
		accumulate512AVX2(acc, input, secret, accWidth)
	}
}

//go:noescape
func scrambleAccAVX2(acc, secret unsafe.Pointer)

func scrambleAcc(acc, secret unsafe.Pointer) {
	switch {
	case useAVX2:
		scrambleAccAVX2(acc, secret)
	}
}

//go:noescape
func accumulateAVX2(acc, input, secret unsafe.Pointer, nbStripes uint64, accWidth ACCWidth)

func accumulate(acc, input, secret unsafe.Pointer, nbStripes uint64, accWidth ACCWidth) {
	switch {
	case useAVX2:
		accumulateAVX2(acc, input, secret, nbStripes, accWidth)
	}
}

func hashLongInternalLoop(acc, input unsafe.Pointer, length uintptr, secret unsafe.Pointer, secretSize uintptr, accWidth ACCWidth) {
	nbRounds := uint64((secretSize - StripeLen) / SecretConsumeRate)
	blockLen := uint64(StripeLen * nbRounds)
	nbBlocks := uint64(length) / blockLen

	if secretSize < SecretSizeMin {
		panic("secretSize is too large than SecretSizeMin")
	}

	for i := uintptr(0); i < uintptr(nbBlocks); i++ {
		accumulate(acc, unsafe.Pointer(uintptr(input)+i*uintptr(blockLen)), secret, nbRounds, accWidth)
		scrambleAcc(acc, unsafe.Pointer(uintptr(secret)+secretSize-uintptr(StripeLen)))
	}

	// last partial block
	if length < StripeLen {
		panic("length is too small than StripeLen")
	}
	nbStripes := uint64((length - uintptr(blockLen*nbBlocks)) / StripeLen)
	if uintptr(nbStripes) < (secretSize / SecretConsumeRate) {
		panic("nbStripes is too small than (secretSize/SecretConsumeRate)")
	}
	accumulate(acc, unsafe.Pointer(uintptr(input)+uintptr(nbBlocks*blockLen)), secret, nbStripes, accWidth)

	// last stripe
	if length&uintptr(StripeLen-1) != 0 {
		p := unsafe.Pointer(uintptr(input) + length - StripeLen)
		const secretLastACCStart = 7
		accumulate512(acc, p, unsafe.Pointer(uintptr(secret)+secretSize-StripeLen-secretLastACCStart), accWidth)
	}
}
