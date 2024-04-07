package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestAssertEquals(t *testing.T) {
	asserter := pkg.NewIntAsserter(t, 10)
	asserter.IsEqualTo(10)
}

func TestAssertFatal(t *testing.T) {
	fT := &FakeT{}
	asserter := pkg.NewIntAsserter(fT, 10)
	asserter.IsEqualTo(9)
	if !fT.isFatal {
		t.Fatalf("FatalF not called")
	}
	if fT.error != "expected 9, but got 10" {
		t.Fatalf("incorrect message = '%v'", fT.error)
	}
}
