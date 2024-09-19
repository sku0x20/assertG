package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

//goland:noinspection GoSnakeCaseUsage
func init_String(t *testing.T) *assertion.Soft {
	return assertion.NewSoft(t)
}

func Test_String(t *testing.T) {
	r := runner.NewTestsRunner(t, init_String)
	r.Add(equalPass_String)
	r.Add(equalFail_String)
	r.Add(containsPass_String)
	r.Add(containsFail_String)
	r.Add(notContainsPass_String)
	r.Add(notContainsFail_String)
	r.Add(hasLengthPass_String)
	r.Add(hasLengthFail_String)
	r.Run()
}

//goland:noinspection GoSnakeCaseUsage
func equalPass_String(t *testing.T, s *assertion.Soft) {
	a := asserter.NewString(s, "some val")
	a.IsEqualTo("some val")
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func equalFail_String(t *testing.T, s *assertion.Soft) {
	a := asserter.NewString(s, "some val")
	a.IsEqualTo("other val")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func containsPass_String(t *testing.T, s *assertion.Soft) {
	a := asserter.NewString(s, "some val")
	a.Contains("val")
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func containsFail_String(t *testing.T, s *assertion.Soft) {
	a := asserter.NewString(s, "some val")
	a.Contains("other")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func notContainsPass_String(t *testing.T, s *assertion.Soft) {
	a := asserter.NewString(s, "some val")
	a.NotContains("other")
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func notContainsFail_String(t *testing.T, s *assertion.Soft) {
	a := asserter.NewString(s, "some val")
	a.NotContains("some")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func hasLengthPass_String(t *testing.T, s *assertion.Soft) {
	a := asserter.NewString(s, "some val")
	a.HasLength(8)
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

//goland:noinspection GoSnakeCaseUsage
func hasLengthFail_String(t *testing.T, s *assertion.Soft) {
	a := asserter.NewString(s, "some val")
	a.HasLength(10)
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}
