package bool

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func init_B(t *testing.T) *assert_type.SoftAssert {
	return assert_type.NewSoftAssert(t)
}

func Test_Bool(t *testing.T) {
	r := runner.NewTestsRunner(t, init_B)
	r.Add(truePass)
	r.Add(trueFail)
	r.Add(falsePass)
	r.Add(falseFail)
	r.Run()
}

func truePass(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewBool(s, 1 == 1)
	a = a.IsTrue()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func trueFail(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewBool(s, 2 == 1)
	a = a.IsTrue()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func falsePass(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewBool(s, 2 == 1)
	a = a.IsFalse()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func falseFail(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewBool(s, 1 == 1)
	a = a.IsFalse()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}
