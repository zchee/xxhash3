// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu_test

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	. "github.com/zchee/xxhash3/internal/cpu"
)

func TestMinimalFeatures(t *testing.T) {
	if runtime.GOARCH == "arm64" {
		switch runtime.GOOS {
		case "linux", "android":
		default:
			t.Skipf("%s/%s is not supported", runtime.GOOS, runtime.GOARCH)
		}
	}

	for _, o := range Options {
		if o.Required && !*o.Feature {
			t.Errorf("%v expected true, got false", o.Name)
		}
	}
}

func MustHaveDebugOptionsSupport(t *testing.T) {
	if !DebugOptions {
		t.Skipf("skipping test: cpu feature options not supported by OS")
	}
}

// hasExec reports whether the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
func HasExec() bool {
	switch runtime.GOOS {
	case "js":
		return false
	case "darwin":
		if strings.HasPrefix(runtime.GOARCH, "arm") {
			return false
		}
	}
	return true
}

// MustHaveExec checks that the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExec calls t.Skip with an explanation.
func MustHaveExec(tb testing.TB) {
	if !HasExec() {
		tb.Skipf("skipping test: cannot exec subprocess on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
}

func runDebugOptionsTest(t *testing.T, test string, options string) {
	MustHaveDebugOptionsSupport(t)

	MustHaveExec(t)

	env := "GODEBUG=" + options

	cmd := exec.Command(os.Args[0], "-test.run="+test)
	cmd.Env = append(cmd.Env, env)

	output, err := cmd.CombinedOutput()
	lines := strings.Fields(string(output))
	lastline := lines[len(lines)-1]

	got := strings.TrimSpace(lastline)
	want := "PASS"
	if err != nil || got != want {
		t.Fatalf("%s with %s: want %s, got %v", test, env, want, got)
	}
}

func TestDisableAllCapabilities(t *testing.T) {
	runDebugOptionsTest(t, "TestAllCapabilitiesDisabled", "cpu.all=off")
}

func TestAllCapabilitiesDisabled(t *testing.T) {
	MustHaveDebugOptionsSupport(t)

	if os.Getenv("GODEBUG") != "cpu.all=off" {
		t.Skipf("skipping test: GODEBUG=cpu.all=off not set")
	}

	for _, o := range Options {
		want := o.Required
		if got := *o.Feature; got != want {
			t.Errorf("%v: expected %v, got %v", o.Name, want, got)
		}
	}
}
