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
	ft.reset()
	ga.IsEqualTo(10)
	ft.assertNotFatal(t)
}

func TestAnyEqualsCallsEqualsIfPossible(t *testing.T) {
	ft := &FakeT{}
	fE1 := &FakeEquals{false}
	fE2 := &FakeEquals{true}
	ga1 := pkg.NewAnyAsserter(ft, fE1)
	ga1.IsEqualTo(fE2)
	ft.assertIsFatal(t, "expected &{true}, got &{false}")
	ft.reset()
	ga2 := pkg.NewAnyAsserter(ft, fE2)
	ga2.IsEqualTo(fE1)
	ft.assertNotFatal(t)
}

func TestAnyNotEquals(t *testing.T) {
	ft := &FakeT{}
	ga := pkg.NewAnyAsserter(ft, 10)
	ga.IsNotEqualTo(10)
	ft.assertIsFatal(t, "expected not 10, got 10")
	ft.reset()
	ga.IsNotEqualTo(20)
	ft.assertNotFatal(t)
}

func TestAnyNotEqualsCallsEqualsIfPossible(t *testing.T) {
	ft := &FakeT{}
	fE1 := &FakeEquals{false}
	fE2 := &FakeEquals{true}
	ga1 := pkg.NewAnyAsserter(ft, fE2)
	ga1.IsNotEqualTo(fE1)
	ft.assertIsFatal(t, "expected not &{false}, got &{true}")
	ft.reset()
	ga2 := pkg.NewAnyAsserter(ft, fE1)
	ga2.IsNotEqualTo(fE2)
	ft.assertNotFatal(t)
}

func TestAnyNil(t *testing.T) {
	ft := &FakeT{}
	actual := 10
	pkg.NewAnyAsserter(ft, &actual).IsNil()
	ft.assertIsFatal(t, fmt.Sprintf("expected <nil>, got %v", &actual))
	ft.reset()
	pkg.NewAnyAsserter(ft, nil).IsNil()
	ft.assertNotFatal(t)
}

func TestAnyNotNil(t *testing.T) {
	ft := &FakeT{}
	pkg.NewAnyAsserter(ft, nil).IsNotNil()
	ft.assertIsFatal(t, "expected not <nil>, got <nil>")
	ft.reset()
	actual := 10
	pkg.NewAnyAsserter(ft, &actual).IsNotNil()
	ft.assertNotFatal(t)
}
