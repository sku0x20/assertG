package asserter

import (
	"assertG/src/pkg/asserter"
	"assertG/src/test/types"
	"testing"
)

func TestAnyEquals(t *testing.T) {
	ft := types.NewFakeT(t)
	ga := asserter.NewAnyAsserter(ft, 10)
	ga.IsEqualTo(30)
	ft.AssertIsFatal()
	ft.Reset()
	ga.IsEqualTo(10)
	ft.AssertNotFatal()
}

func TestAnyEqualsCallsEqualsIfPossible(t *testing.T) {
	ft := types.NewFakeT(t)
	equalFalse := &types.FakeEquals{false}
	equalTrue := &types.FakeEquals{true}
	ga1 := asserter.NewAnyAsserter(ft, equalFalse)
	ga1.IsEqualTo(equalTrue)
	ft.AssertIsFatal()
	ft.Reset()
	ga2 := asserter.NewAnyAsserter(ft, equalTrue)
	ga2.IsEqualTo(equalFalse)
	ft.AssertNotFatal()
}

func TestAnyNotEquals(t *testing.T) {
	ft := types.NewFakeT(t)
	ga := asserter.NewAnyAsserter(ft, 10)
	ga.IsNotEqualTo(10)
	ft.AssertIsFatal()
	ft.Reset()
	ga.IsNotEqualTo(20)
	ft.AssertNotFatal()
}

func TestAnyNotEqualsCallsEqualsIfPossible(t *testing.T) {
	ft := types.NewFakeT(t)
	fE1 := &types.FakeEquals{false}
	fE2 := &types.FakeEquals{true}
	ga1 := asserter.NewAnyAsserter(ft, fE2)
	ga1.IsNotEqualTo(fE1)
	ft.AssertIsFatal()
	ft.Reset()
	ga2 := asserter.NewAnyAsserter(ft, fE1)
	ga2.IsNotEqualTo(fE2)
	ft.AssertNotFatal()
}

func TestAnyNil(t *testing.T) {
	ft := types.NewFakeT(t)
	actual := 10
	asserter.NewAnyAsserter(ft, &actual).IsNil()
	ft.AssertIsFatal()
	ft.Reset()
	asserter.NewAnyAsserter(ft, nil).IsNil()
	ft.AssertNotFatal()
}

func TestAnyNotNil(t *testing.T) {
	ft := types.NewFakeT(t)
	asserter.NewAnyAsserter(ft, nil).IsNotNil()
	ft.AssertIsFatal()
	ft.Reset()
	actual := 10
	asserter.NewAnyAsserter(ft, &actual).IsNotNil()
	ft.AssertNotFatal()
}
