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
	r.Add(ContainsExactlyAnyOrder_Slice)
	r.Run()
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

type s struct {
	a int
	b int
}
type test struct {
	a    []s
	fail bool
}

func ContainsExactlyInOrder_Slice(tm *testing.T, _ *assertion.Soft) {
	e := []s{{10, 20}, {20, 30}}
	tests := []test{
		{[]s{{10, 20}, {20, 30}}, false},
		{[]s{{10, 20}, {20, 10}}, true},
		{[]s{{10, 20}}, true},
		{[]s{{20, 30}, {10, 20}}, true},
	}
	for _, ct := range tests {
		tm.Run("", func(t *testing.T) {
			sa := assertion.NewSoft(t)
			a := asserter.NewSlice(sa, e)
			a = a.Contains(asserter.EXACTLY, asserter.IN_ORDER, ct.a...)
			if sa.Failed() != ct.fail {
				t.Fatalf("should have failed = %t", ct.fail)
			}
		})
	}
}
