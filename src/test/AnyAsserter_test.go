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

func TestAnyNil(t *testing.T) {
	ft := &FakeT{}
	actual := 10
	ga := pkg.NewAnyAsserter(ft, &actual)
	ga.IsNil()
	ft.assertIsFatal(t, fmt.Sprintf("expected <nil>, got %v", &actual))
}
