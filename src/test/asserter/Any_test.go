package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/test/api"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

//goland:noinspection GoSnakeCaseUsage
func init_Any(t *testing.T) *assertion.Soft {
	return assertion.NewSoft(t)
}

func Test_Any(t *testing.T) {
	r := runner.NewTestsRunner[*assertion.Soft](t, init_Any)
	r.Add(notEqual_Any)
	r.Add(equal_Any)
	r.Add(implementsEqual_Any)
	r.Add(nilPass_Any)
	r.Add(nilFail_Any)
	r.Add(notNilPass_Any)
	r.Add(notNilFail_Any)
	r.Run()
}

//goland:noinspection GoSnakeCaseUsage
func notEqual_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, 30)
	a.IsEqualTo("something")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func equal_Any(t *testing.T, s *assertion.Soft) {
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

//goland:noinspection GoSnakeCaseUsage
func implementsEqual_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, api.NewFakeEqual(false))
	a.IsEqualTo(api.NewFakeEqual(false))
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func nilPass_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, nil)
	a.IsNil()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func nilFail_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, 10)
	a.IsNil()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func notNilPass_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, 10)
	a.IsNotNil()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func notNilFail_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, nil)
	a.IsNotNil()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}
