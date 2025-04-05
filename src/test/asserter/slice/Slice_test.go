package slice

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Slice(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(hasLength)
	//r.Add(ContainsExactlyInOrder_Slice)
	//r.Add(ContainsExactlyAnyOrder_Slice)
	r.Run()
}

func hasLength(t *testing.T, _ any) {
	intS := []int{10}
	sa := assertion.NewSoft(t)
	a := asserter.NewSlice(sa, intS)
	a = a.HasLength(1)
	if sa.Failed() {
		t.Fatalf("should not have failed")
	}
	sa = assertion.NewSoft(t)
	a = asserter.NewSlice(sa, intS)
	a = a.HasLength(2)
	if !sa.Failed() {
		t.Fatalf("should have failed")
	}
}

type s struct {
	a int
	b int
}
type test struct {
	e    []s
	a    []s
	fail bool
}

func ContainsExactlyInOrder_Slice(tm *testing.T, _ *assertion.Soft) {
	e := []s{{10, 20}, {20, 30}}
	tests := []test{
		{e, []s{{10, 20}, {20, 30}}, false},
		{e, []s{{10, 20}, {20, 10}}, true},
		{e, []s{{10, 20}}, true},
		{e, []s{{20, 30}, {10, 20}}, true},
	}
	runTests(tm, tests, func(a *asserter.Slice[s], elem ...s) *asserter.Slice[s] {
		return a.Contains(asserter.EXACTLY, asserter.IN_ORDER, elem...)
	})
}

func ContainsExactlyAnyOrder_Slice(tm *testing.T, _ *assertion.Soft) {
	e := []s{{10, 20}, {20, 30}}
	tests := []test{
		{e, []s{{10, 20}, {20, 30}}, false},
		{e, []s{{10, 20}, {20, 10}}, true},
		{e, []s{{10, 20}}, true},
		{e, []s{{20, 30}, {10, 20}}, false},
	}
	runTests(tm, tests, func(a *asserter.Slice[s], elem ...s) *asserter.Slice[s] {
		return a.Contains(asserter.EXACTLY, asserter.ANY_ORDER, elem...)
	})
}

func runTests(
	t *testing.T,
	tests []test,
	fnAssertion func(a *asserter.Slice[s], elem ...s) *asserter.Slice[s],
) {
	for _, ct := range tests {
		t.Run("", func(t *testing.T) {
			sa := assertion.NewSoft(t)
			a := asserter.NewSlice(sa, ct.e)
			a = fnAssertion(a, ct.a...)
			if sa.Failed() != ct.fail {
				t.Fatalf("should have failed = %t", ct.fail)
			}
		})
	}
}
