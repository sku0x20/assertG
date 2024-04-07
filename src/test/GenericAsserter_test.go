package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestGenericEquals(t *testing.T) {
	ft := &FakeT{}
	ga := pkg.NewGenericAsserter(ft, 10)
	ga.IsEqualTo(30)
	if !ft.isFatal {
		t.Fatalf("should be fatal")
	}
	if ft.error != "expected 30, got 10" {
		t.Fatalf("invalid message = '%v'", ft.error)
	}
}
