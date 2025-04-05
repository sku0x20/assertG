package string

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func init_S(t *testing.T) *assert_type.SoftAssert {
	return assert_type.NewSoftAssert(t)
}

func Test_String(t *testing.T) {
	r := runner.NewTestsRunner(t, init_S)
	r.Add(equalPass)
	r.Add(equalFail)
	r.Add(containsPass)
	r.Add(containsFail)
	r.Add(notContainsPass)
	r.Add(notContainsFail)
	r.Add(hasLengthPass)
	r.Add(hasLengthFail)
	r.Run()
}

func equalPass(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewString(s, "some val")
	a = a.IsEqualTo("some val")
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func equalFail(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewString(s, "some val")
	a = a.IsEqualTo("other val")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func containsPass(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewString(s, "some val")
	a = a.Contains("val")
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func containsFail(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewString(s, "some val")
	a = a.Contains("other")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func notContainsPass(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewString(s, "some val")
	a = a.NotContains("other")
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func notContainsFail(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewString(s, "some val")
	a = a.NotContains("some")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func hasLengthPass(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewString(s, "some val")
	a = a.HasLength(8)
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func hasLengthFail(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewString(s, "some val")
	a = a.HasLength(10)
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}
