package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func init_Bool(t *testing.T) *assertion.Soft {
	return assertion.NewSoft(t)
}

func Test_Bool(t *testing.T) {
	r := runner.NewTestsRunner(t, init_Bool)
	r.Add(truePass_Bool)
	r.Add(trueFail_Bool)
	r.Add(falsePass_Bool)
	r.Add(falseFail_Bool)
	r.Run()
}

func truePass_Bool(t *testing.T, s *assertion.Soft) {
	a := asserter.NewBool(s, 1 == 1)
	a = a.IsTrue()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func trueFail_Bool(t *testing.T, s *assertion.Soft) {
	a := asserter.NewBool(s, 2 == 1)
	a = a.IsTrue()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func falsePass_Bool(t *testing.T, s *assertion.Soft) {
	a := asserter.NewBool(s, 2 == 1)
	a = a.IsFalse()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func falseFail_Bool(t *testing.T, s *assertion.Soft) {
	a := asserter.NewBool(s, 1 == 1)
	a = a.IsFalse()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}
