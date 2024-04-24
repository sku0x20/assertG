package asserter

import (
	"assertG/src/pkg/asserter"
	"assertG/src/test/asserter/fake"
	"testing"
)

func TestStringEquals(t *testing.T) {
	ft := fake.NewFakeT(t)
	sa := asserter.NewStringAsserter(ft, "some-string")
	sa = sa.IsEqualTo("other-string")
	ft.AssertIsFatal()
	ft.Reset()
	sa.IsEqualTo("some-string")
	ft.AssertNotFatal()
}

func TestStringNotEquals(t *testing.T) {
	ft := fake.NewFakeT(t)
	sa := asserter.NewStringAsserter(ft, "some-string")
	sa = sa.IsNotEqualTo("some-string")
	ft.AssertIsFatal()
	ft.Reset()
	sa.IsNotEqualTo("other-string")
	ft.AssertNotFatal()
}

func TestContains(t *testing.T) {
	ft := fake.NewFakeT(t)
	sa := asserter.NewStringAsserter(ft, "some-string")
	sa = sa.Contains("other")
	ft.AssertIsFatal()
	ft.Reset()
	sa.Contains("some")
	ft.AssertNotFatal()
}

func TestNotContains(t *testing.T) {
	ft := fake.NewFakeT(t)
	sa := asserter.NewStringAsserter(ft, "some-string")
	sa = sa.NotContains("some")
	ft.AssertIsFatal()
	ft.Reset()
	sa.NotContains("other")
	ft.AssertNotFatal()
}

func TestLength(t *testing.T) {
	ft := fake.NewFakeT(t)
	sa := asserter.NewStringAsserter(ft, "some-string")
	sa = sa.HasLength(8)
	ft.AssertIsFatal()
	ft.Reset()
	sa.HasLength(11)
	ft.AssertNotFatal()
}
