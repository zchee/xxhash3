// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

var CPUID func(op uint32) (eax, ebx, ecx, edx uint32)
var CPUIndex func(op, op2 uint32) (eax, ebx, ecx, edx uint32)
var Xgetbv func(index uint32) (eax, edx uint32)
var Rdtscp func() (eax, ebx, ecx, edx uint32)

// DebugOptions is set to true by the runtime if the OS supports reading
// GODEBUG early in runtime startup.
// This should not be changed after it is initialized.
var DebugOptions bool

// CacheLinePad is used to pad structs to avoid false sharing.
type CacheLinePad struct{ _ [CacheLinePadSize]byte }

// CacheLineSize is the CPU's assumed cache line size.
// There is currently no runtime detection of the real cache line size
// so we use the constant per GOARCH CacheLinePadSize as an approximation.
var CacheLineSize uintptr = CacheLinePadSize

// X86 contains the supported CPU features of the
// current X86/AMD64 platform. If the current platform
// is not X86/AMD64 then all feature flags are false.
//
// X86 is padded to avoid false sharing. Further the HasAVX
// and HasAVX2 are only set if the OS supports XMM and YMM
// registers in addition to the CPUID feature bit being set.
var X86 x86

type x86 struct {
	_            CacheLinePad
	HasAES       bool
	HasADX       bool
	HasAVX       bool
	HasAVX2      bool
	HasBMI1      bool
	HasBMI2      bool
	HasERMS      bool
	HasFMA       bool
	HasOSXSAVE   bool
	HasPCLMULQDQ bool
	HasPOPCNT    bool
	HasSSE2      bool
	HasSSE3      bool
	HasSSSE3     bool
	HasSSE41     bool
	HasSSE42     bool
	_            CacheLinePad
}

// ARM64 contains the supported CPU features of the
// current ARMv8(aarch64) platform. If the current platform
// is not arm64 then all feature flags are false.
var ARM64 arm64

type arm64 struct {
	_           CacheLinePad
	HasFP       bool
	HasASIMD    bool
	HasEVTSTRM  bool
	HasAES      bool
	HasPMULL    bool
	HasSHA1     bool
	HasSHA2     bool
	HasCRC32    bool
	HasATOMICS  bool
	HasFPHP     bool
	HasASIMDHP  bool
	HasCPUID    bool
	HasASIMDRDM bool
	HasJSCVT    bool
	HasFCMA     bool
	HasLRCPC    bool
	HasDCPOP    bool
	HasSHA3     bool
	HasSM3      bool
	HasSM4      bool
	HasASIMDDP  bool
	HasSHA512   bool
	HasSVE      bool
	HasASIMDFHM bool
	_           CacheLinePad
}

// ARM contains the supported CPU features of the current ARM (32-bit) platform.
// All feature flags are false if:
//   1. the current platform is not arm, or
//   2. the current operating system is not Linux.
var ARM arm

type arm struct {
	_           CacheLinePad
	HasSWP      bool // SWP instruction support
	HasHALF     bool // Half-word load and store support
	HasTHUMB    bool // ARM Thumb instruction set
	Has26BIT    bool // Address space limited to 26-bits
	HasFASTMUL  bool // 32-bit operand, 64-bit result multiplication support
	HasFPA      bool // Floating point arithmetic support
	HasVFP      bool // Vector floating point support
	HasEDSP     bool // DSP Extensions support
	HasJAVA     bool // Java instruction set
	HasIWMMXT   bool // Intel Wireless MMX technology support
	HasCRUNCH   bool // MaverickCrunch context switching and handling
	HasTHUMBEE  bool // Thumb EE instruction set
	HasNEON     bool // NEON instruction set
	HasVFPv3    bool // Vector floating point version 3 support
	HasVFPv3D16 bool // Vector floating point version 3 D8-D15
	HasTLS      bool // Thread local storage support
	HasVFPv4    bool // Vector floating point version 4 support
	HasIDIVA    bool // Integer divide instruction support in ARM mode
	HasIDIVT    bool // Integer divide instruction support in Thumb mode
	HasVFPD32   bool // Vector floating point version 3 D15-D31
	HasLPAE     bool // Large Physical Address Extensions
	HasEVTSTRM  bool // Event stream support
	HasAES      bool // AES hardware implementation
	HasPMULL    bool // Polynomial multiplication instruction set
	HasSHA1     bool // SHA1 hardware implementation
	HasSHA2     bool // SHA2 hardware implementation
	HasCRC32    bool // CRC32 hardware implementation
	_           CacheLinePad
}

// PPC64 contains the supported CPU features of the current ppc64/ppc64le platforms.
// If the current platform is not ppc64/ppc64le then all feature flags are false.
//
// For ppc64/ppc64le, it is safe to check only for ISA level starting on ISA v3.00,
// since there are no optional categories. There are some exceptions that also
// require kernel support to work (DARN, SCV), so there are feature bits for
// those as well. The minimum processor requirement is POWER8 (ISA 2.07).
// The struct is padded to avoid false sharing.
var PPC64 ppc64

type ppc64 struct {
	_        CacheLinePad
	HasDARN  bool // Hardware random number generator (requires kernel enablement)
	HasSCV   bool // Syscall vectored (requires kernel enablement)
	IsPOWER8 bool // ISA v2.07 (POWER8)
	IsPOWER9 bool // ISA v3.00 (POWER9)
	_        CacheLinePad
}

// S390X contains the supported CPU features of the current IBM Z
// (s390x) platform. If the current platform is not IBM Z then all
// feature flags are false.
//
// S390X is padded to avoid false sharing. Further HasVX is only set
// if the OS supports vector registers in addition to the STFLE
// feature bit being set.
var S390X s390x

type s390x struct {
	_         CacheLinePad
	HasZARCH  bool // z architecture mode is active [mandatory]
	HasSTFLE  bool // store facility list extended [mandatory]
	HasLDISP  bool // long (20-bit) displacements [mandatory]
	HasEIMM   bool // 32-bit immediates [mandatory]
	HasDFP    bool // decimal floating point
	HasETF3EH bool // ETF-3 enhanced
	HasMSA    bool // message security assist (CPACF)
	HasAES    bool // KM-AES{128,192,256} functions
	HasAESCBC bool // KMC-AES{128,192,256} functions
	HasAESCTR bool // KMCTR-AES{128,192,256} functions
	HasAESGCM bool // KMA-GCM-AES{128,192,256} functions
	HasGHASH  bool // KIMD-GHASH function
	HasSHA1   bool // K{I,L}MD-SHA-1 functions
	HasSHA256 bool // K{I,L}MD-SHA-256 functions
	HasSHA512 bool // K{I,L}MD-SHA-512 functions
	HasSHA3   bool // K{I,L}MD-SHA3-{224,256,384,512} and K{I,L}MD-SHAKE-{128,256} functions
	HasVX     bool // vector facility. Note: the runtime sets this when it processes auxv records.
	HasVXE    bool // vector-enhancements facility 1
	HasKDSA   bool // elliptic curve functions
	HasECDSA  bool // NIST curves
	HasEDDSA  bool // Edwards curves
	_         CacheLinePad
}
