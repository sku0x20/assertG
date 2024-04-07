package test

import (
	"assertG/src/pkg"
	"fmt"
	"testing"
)

func TestAnyEquals(t *testing.T) {
	ft := &FakeT{}
	ga := pkg.NewAnyAsserter(ft, 10)
	ga.IsEqualTo(30)
	ft.assertIsFatal(t, "expected 30, got 10")
}

func TestAnyNotEquals(t *testing.T) {
	ft := &FakeT{}
	ga := pkg.NewAnyAsserter(ft, 10)
	ga.IsNotEqualTo(10)
	ft.assertIsFatal(t, "expected not 10, got 10")
}

func TestAnyNil(t *testing.T) {
	ft := &FakeT{}
	actual := 10
	ga := pkg.NewAnyAsserter(ft, &actual)
	ga.IsNil()
	ft.assertIsFatal(t, fmt.Sprintf("expected <nil>, got %v", &actual))
}

func TestAnyNotNil(t *testing.T) {
	ft := &FakeT{}
	ga := pkg.NewAnyAsserter(ft, nil)
	ga.IsNotNil()
	ft.assertIsFatal(t, "expected not <nil>, got <nil>")
}
