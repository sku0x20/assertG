package test

import (
	"assertG/src/pkg"
	"fmt"
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

func TestGenericNil(t *testing.T) {
	ft := &FakeT{}
	actual := 10
	ga := pkg.NewGenericAsserter(ft, &actual)
	ga.IsNil()
	if !ft.isFatal {
		t.Fatalf("should be fatal")
	}
	if ft.error != fmt.Sprintf("expected <nil>, got %v", &actual) {
		t.Fatalf("invalid message = '%v'", ft.error)
	}
}
