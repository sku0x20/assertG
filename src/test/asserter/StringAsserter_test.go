package asserter

import (
	"assertG/src/pkg"
	"assertG/src/pkg/asserter"
	"assertG/src/test/types"
	"testing"
)

func TestStringEquals(t *testing.T) {
	ft := types.NewFakeT(t)
	h := pkg.NewAssertHelper(ft, "some-string")
	sa := asserter.NewStringAsserter(h)
	sa = sa.IsEqualTo("other-string")
	ft.AssertIsFatal()
	ft.Reset()
	sa.IsEqualTo("some-string")
	ft.AssertNotFatal()
}

func TestStringNotEquals(t *testing.T) {
	ft := types.NewFakeT(t)
	h := pkg.NewAssertHelper(ft, "some-string")
	sa := asserter.NewStringAsserter(h)
	sa = sa.IsNotEqualTo("some-string")
	ft.AssertIsFatal()
	ft.Reset()
	sa.IsNotEqualTo("other-string")
	ft.AssertNotFatal()
}

func TestContains(t *testing.T) {
	ft := types.NewFakeT(t)
	h := pkg.NewAssertHelper(ft, "some-string")
	sa := asserter.NewStringAsserter(h)
	sa = sa.Contains("other")
	ft.AssertIsFatal()
	ft.Reset()
	sa.Contains("some")
	ft.AssertNotFatal()
}

func TestNotContains(t *testing.T) {
	ft := types.NewFakeT(t)
	h := pkg.NewAssertHelper(ft, "some-string")
	sa := asserter.NewStringAsserter(h)
	sa = sa.NotContains("some")
	ft.AssertIsFatal()
	ft.Reset()
	sa.NotContains("other")
	ft.AssertNotFatal()
}

func TestLength(t *testing.T) {
	ft := types.NewFakeT(t)
	h := pkg.NewAssertHelper(ft, "some-string")
	sa := asserter.NewStringAsserter(h)
	sa = sa.HasLength(8)
	ft.AssertIsFatal()
	ft.Reset()
	sa.HasLength(11)
	ft.AssertNotFatal()
}
