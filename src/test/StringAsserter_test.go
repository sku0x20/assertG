package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestStringEquals(t *testing.T) {
	ft := &FakeT{}
	sa := pkg.NewStringAsserter(ft, "some-string")
	sa = sa.IsEqualTo("other-string")
	ft.assertIsFatal(t, "expected 'other-string', got 'some-string'")
	ft.reset()
	sa.IsEqualTo("some-string")
	ft.assertNotFatal(t)
}

func TestStringNotEquals(t *testing.T) {
	ft := &FakeT{}
	sa := pkg.NewStringAsserter(ft, "some-string")
	sa = sa.IsNotEqualTo("some-string")
	ft.assertIsFatal(t, "not expected 'some-string', got 'some-string'")
	ft.reset()
	sa.IsNotEqualTo("other-string")
	ft.assertNotFatal(t)
}

func TestContains(t *testing.T) {
	ft := &FakeT{}
	sa := pkg.NewStringAsserter(ft, "some-string")
	sa = sa.Contains("other")
	ft.assertIsFatal(t, "expected contains 'other', got 'some-string'")
	ft.reset()
	sa.Contains("some")
	ft.assertNotFatal(t)
}

func TestNotContains(t *testing.T) {
	ft := &FakeT{}
	sa := pkg.NewStringAsserter(ft, "some-string")
	sa = sa.NotContains("some")
	ft.assertIsFatal(t, "expected not contains 'some', got 'some-string'")
	ft.reset()
	sa.NotContains("other")
	ft.assertNotFatal(t)
}
