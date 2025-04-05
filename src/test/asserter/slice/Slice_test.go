package slice

import (
	"github.com/sku0x20/assertG/src/pkg/assert_type"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Slice(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(hasLength)
	r.Add(IsEqualTo)
	//r.Add(ContainsExactlyAnyOrder_Slice)
	r.Run()
}

func hasLength(t *testing.T, _ any) {
	intS := []int{10}
	sa := assert_type.NewSoftAssert(t)
	a := asserter.NewSlice(sa, intS)
	a = a.HasLength(1)
	if sa.Failed() {
		t.Fatalf("should not have failed")
	}
	sa = assert_type.NewSoftAssert(t)
	a = asserter.NewSlice(sa, intS)
	a = a.HasLength(2)
	if !sa.Failed() {
		t.Fatalf("should have failed")
	}
}

type test struct {
	actual     []int
	expected   []int
	shouldFail bool
}

func IsEqualTo(t *testing.T, _ any) {
	tests := []test{
		{[]int{10, 20, 30}, []int{10}, true},
		{[]int{10, 20, 30}, []int{10, 20, 30, 40}, true},
		{[]int{10, 20, 30}, []int{30, 10, 20}, true},
		{[]int{10, 20, 30}, []int{10, 20, 30}, false},
	}
	runTests(t, tests, func(a *asserter.Slice[int], expected []int) *asserter.Slice[int] {
		return a.IsEqualTo(expected)
	})
}

func runTests(
	t *testing.T,
	tests []test,
	fnAssertion func(a *asserter.Slice[int], elem []int) *asserter.Slice[int],
) {
	for _, ct := range tests {
		t.Run("", func(t *testing.T) {
			sa := assert_type.NewSoftAssert(t)
			a := asserter.NewSlice(sa, ct.actual)
			a = fnAssertion(a, ct.expected)
			if sa.Failed() != ct.shouldFail {
				t.Fatalf("should have failed")
			}
		})
	}
}
