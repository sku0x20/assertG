package asserter

import (
	"assertG/src/pkg"
	"assertG/src/pkg/asserter"
	"assertG/src/test/asserter/fake"
	"testing"
)

func TestStringEquals(t *testing.T) {
	ft := &fake.FakeT{}
	sa := asserter.NewStringAsserter(ft, pkg.NewFormatter(), "some-string")
	sa = sa.IsEqualTo("other-string")
	ft.AssertIsFatal(t)
	ft.Reset()
	sa.IsEqualTo("some-string")
	ft.AssertNotFatal(t)
}

func TestStringNotEquals(t *testing.T) {
	ft := &fake.FakeT{}
	sa := asserter.NewStringAsserter(ft, pkg.NewFormatter(), "some-string")
	sa = sa.IsNotEqualTo("some-string")
	ft.AssertIsFatal(t)
	ft.Reset()
	sa.IsNotEqualTo("other-string")
	ft.AssertNotFatal(t)
}

func TestContains(t *testing.T) {
	ft := &fake.FakeT{}
	sa := asserter.NewStringAsserter(ft, pkg.NewFormatter(), "some-string")
	sa = sa.Contains("other")
	ft.AssertIsFatal(t)
	ft.Reset()
	sa.Contains("some")
	ft.AssertNotFatal(t)
}

func TestNotContains(t *testing.T) {
	ft := &fake.FakeT{}
	sa := asserter.NewStringAsserter(ft, pkg.NewFormatter(), "some-string")
	sa = sa.NotContains("some")
	ft.AssertIsFatal(t)
	ft.Reset()
	sa.NotContains("other")
	ft.AssertNotFatal(t)
}

func TestLength(t *testing.T) {
	ft := &fake.FakeT{}
	sa := asserter.NewStringAsserter(ft, pkg.NewFormatter(), "some-string")
	sa = sa.HasLength(8)
	ft.AssertIsFatal(t)
	ft.Reset()
	sa.HasLength(11)
	ft.AssertNotFatal(t)
}
