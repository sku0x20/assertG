package test

import (
	"assertG/src/pkg"
	"fmt"
	"testing"
)

func TestAssertEquals(t *testing.T) {
	asserter := pkg.NewIntAsserter(t, 10)
	asserter.IsEqualTo(10)
}

func TestAssertFatal(t *testing.T) {
	fT := &fakeT{}
	asserter := pkg.NewIntAsserter(fT, 10)
	asserter.IsEqualTo(9)
	if !fT.isCalled {
		t.Fatalf("FatalF not called")
	}
	if fT.error != "expected 9, but got 10" {
		t.Fatalf("incorrect message = '%v'", fT.error)
	}
}

type fakeT struct {
	isCalled bool
	error    string
}

func (t *fakeT) Fatalf(format string, args ...any) {
	t.isCalled = true
	t.error = fmt.Sprintf(format, args...)
}
