package slice

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/equator"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Slice(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(hasLength)
	r.Add(IsEqualTo)
	//r.Add(IsEqualToIgnoringOrder)
	r.Run()
}

func hasLength(t *testing.T, _ any) {
	intS := []int{10}
	sa := assert_type.NewSoftAssert(t)
	a := asserter.NewSlice(sa, equator.NewComparableEquator[int](), intS)
	a = a.HasLength(1)
	if sa.Failed() {
		t.Fatalf("should not have failed")
	}
	sa = assert_type.NewSoftAssert(t)
	a = asserter.NewSlice(sa, equator.NewComparableEquator[int](), intS)
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
		defer func() {
			if r := recover(); r != nil {
				// panic is valid in the 1st and 2nd case
				t.Logf("recovered from panic: %v", r)
			}
		}()
		return a.IsEqualTo(expected)
	})
}

//func IsEqualToIgnoringOrder(t *testing.T, _ any) {
//	t.Skip()
//	tests := []test{
//		{[]int{10, 20, 30}, []int{10}, true},
//		{[]int{10, 20, 30}, []int{10, 20, 30, 40}, true},
//		{[]int{10, 20, 30}, []int{0, 20, 30}, true},
//		{[]int{10, 20, 30}, []int{30, 10, 20}, false},
//		{[]int{10, 20, 30}, []int{10, 20, 30}, false},
//	}
//	runTests(t, tests, func(a *asserter.Slice[int], expected []int) *asserter.Slice[int] {
//		defer func() {
//			if r := recover(); r != nil {
//				// panic is valid in the 1st and 2nd case
//				t.Logf("recovered from panic: %v", r)
//			}
//		}()
//		return a.IsEqualToIgnoringOrder(expected)
//	})
//}

func runTests(
	t *testing.T,
	tests []test,
	fnAssertion func(a *asserter.Slice[int], elem []int) *asserter.Slice[int],
) {
	for _, ct := range tests {
		t.Run("", func(t *testing.T) {
			sa := assert_type.NewSoftAssert(t)
			a := asserter.NewSlice(sa, equator.NewComparableEquator[int](), ct.actual)
			a = fnAssertion(a, ct.expected)
			if sa.Failed() != ct.shouldFail {
				t.Fatalf("should have failed = %t", ct.shouldFail)
			}
		})
	}
}
