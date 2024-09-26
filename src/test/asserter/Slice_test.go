package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func init_Slice(t *testing.T) *assertion.Soft {
	return assertion.NewSoft(t)
}

func Test_Slice(t *testing.T) {
	r := runner.NewTestsRunner(t, init_Slice)
	r.Add(hasLengthPass_Slice)
	r.Add(hasLengthFail_Slice)
	r.Add(ContainsExactlyInOrder_Slice)
	r.Run()
}

func ContainsExactlyInOrder_Slice(tm *testing.T, _ *assertion.Soft) {
	tests := []struct {
		e    []int
		a    []int
		fail bool
	}{
		{[]int{1, 2}, []int{1, 2}, false},
		{[]int{1, 2}, []int{1, 3}, true},
		{[]int{1, 2}, []int{1}, true},
		{[]int{1, 2}, []int{2, 1}, true},
	}
	for _, test := range tests {
		tm.Run("", func(t *testing.T) {
			s := assertion.NewSoft(t)
			a := asserter.NewSlice(s, test.e)
			a = a.Contains(asserter.EXACTLY, asserter.IN_ORDER, test.a...)
			if s.Failed() != test.fail {
				t.Fatalf("should have failed = %t", test.fail)
			}
		})
	}
}

func hasLengthPass_Slice(t *testing.T, s *assertion.Soft) {
	intS := []int{10}
	a := asserter.NewSlice(s, intS)
	a = a.HasLength(1)
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func hasLengthFail_Slice(t *testing.T, s *assertion.Soft) {
	intS := []int{10}
	a := asserter.NewSlice(s, intS)
	a = a.HasLength(2)
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}
