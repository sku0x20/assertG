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

/*
func TestStringNil(t *testing.T) {
	ft := &FakeT{}
	actual := 10
	pkg.NewAnyAsserter(ft, &actual).IsNil()
	ft.assertIsFatal(t, fmt.Sprintf("expected '<nil>', got '%v'", &actual))
	ft.reset()
	pkg.NewAnyAsserter(ft, nil).IsNil()
	ft.assertNotFatal(t)
}

func TestStringNotNil(t *testing.T) {
	ft := &FakeT{}
	pkg.NewAnyAsserter(ft, nil).IsNotNil()
	ft.assertIsFatal(t, "expected 'not <nil>', got '<nil>'")
	ft.reset()
	actual := 10
	pkg.NewAnyAsserter(ft, &actual).IsNotNil()
	ft.assertNotFatal(t)
}
*/
