package assert

import (
	"assertG/src/pkg"
	. "assertG/src/pkg/assert"
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	code := t.Run()
	println("called code is", code)
	os.Exit(code)
}

func TestAssertThatInt1(t *testing.T) {
	pkg.RegisterT(t)
	AssertThat(12).IsEqualTo(12)
}

func TestAssertThatInt2(t *testing.T) {
	pkg.RegisterT(t)
	AssertThat(12).IsEqualTo(12)
}

// support assert.That(12).isEqualTo(12) also
