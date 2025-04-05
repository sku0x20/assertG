package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/pkg/equator"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func init_Any(t *testing.T) *assertion.Soft {
	return assertion.NewSoft(t)
}

func Test_Any(t *testing.T) {
	r := runner.NewTestsRunner[*assertion.Soft](t, init_Any)
	r.Add(equalFail_Any)
	r.Add(equalPass_Any)
	r.Add(customEquator_Any)
	r.Add(nilPass_Any)
	r.Add(nilFail_Any)
	r.Add(notNilPass_Any)
	r.Add(notNilFail_Any)
	r.Run()
}

func equalFail_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, equator.NewReflectDeepEquator[any](), 30)
	a = a.IsEqualTo("something")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func equalPass_Any(t *testing.T, s *assertion.Soft) {
	type st struct {
		s *string
	}
	s1 := "something"
	s2 := "something"
	a := asserter.NewAny(s, equator.NewReflectDeepEquator[any](), &st{&s1})
	a = a.IsEqualTo(&st{&s2})
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

type FakeEquator struct {
}

func (f *FakeEquator) AreEqual(a any, b any) bool {
	return true
}

func customEquator_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, &FakeEquator{}, 10)
	a = a.IsEqualTo(20)
	if s.Failed() {
		t.Fatalf("should have passed")
	}
}

func nilPass_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, equator.NewReflectDeepEquator[any](), nil)
	a = a.IsNil()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func nilFail_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, equator.NewReflectDeepEquator[any](), 10)
	a = a.IsNil()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func notNilPass_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, equator.NewReflectDeepEquator[any](), 10)
	a = a.IsNotNil()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func notNilFail_Any(t *testing.T, s *assertion.Soft) {
	a := asserter.NewAny(s, equator.NewReflectDeepEquator[any](), nil)
	a = a.IsNotNil()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}
