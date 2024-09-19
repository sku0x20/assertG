package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/test/api"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Any(t *testing.T) {
	r := runner.NewTestsRunner[*assertion.Soft](t, initE)
	r.Add(notEqual)
	r.Add(equal)
	r.Add(implementsEqual)
	r.Run()
}

func initE(t *testing.T) *assertion.Soft {
	return assertion.NewSoft(t)
}

func notEqual(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, 30)
	a.IsEqualTo("something")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func equal(t *testing.T, s *assertion.Soft) {
	type st struct {
		s *string
	}
	s1 := "something"
	s2 := "something"
	a := asserter.NewAny(s, &st{&s1})
	a.IsEqualTo(&st{&s2})
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func implementsEqual(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, api.NewFakeEqual(false))
	a.IsEqualTo(api.NewFakeEqual(false))
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

/*

func TestAnyNil(t *testing.T) {
	ft := api.NewFakeT(t)
	actual := 10
	h := pkg.NewAssertHelper(ft, &actual)
	asserter.NewAnyAsserter(h).IsNil()
	ft.AssertIsFatal()
	ft.Reset()
	h2 := pkg.NewAssertHelper(ft, nil)
	asserter.NewAnyAsserter(h2).IsNil()
	ft.AssertNotFatal()
}

func TestAnyNotNil(t *testing.T) {
	ft := api.NewFakeT(t)
	h := pkg.NewAssertHelper(ft, nil)
	asserter.NewAnyAsserter(h).IsNotNil()
	ft.AssertIsFatal()
	ft.Reset()
	actual := 10
	h2 := pkg.NewAssertHelper(ft, &actual)
	asserter.NewAnyAsserter(h2).IsNotNil()
	ft.AssertNotFatal()
}
*/
