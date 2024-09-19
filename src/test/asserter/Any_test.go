package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/test/api"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func initE(t *testing.T) *assertion.Soft {
	return assertion.NewSoft(t)
}

func Test_Any(t *testing.T) {
	r := runner.NewTestsRunner[*assertion.Soft](t, initE)
	r.Add(notEqual)
	r.Add(equal)
	r.Add(implementsEqual)
	r.Add(nilPass)
	r.Add(nilFail)
	r.Add(notNilPass)
	r.Add(notNilFail)
	r.Run()
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

func nilPass(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, nil)
	a.IsNil()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func nilFail(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, 10)
	a.IsNil()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func notNilPass(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, 10)
	a.IsNotNil()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func notNilFail(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, nil)
	a.IsNotNil()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}
