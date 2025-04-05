package assert

import (
	"github.com/sku0x20/assertG/src/main/assert"
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/equator"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Slice(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(thatSlice)
	r.Add(thatSliceWith)
	r.Run()
}

func thatSlice(t *testing.T, _ any) {
	var a any = assert.ThatSlice([]string{"some val"})
	casted, ok := a.(*asserter.Slice[string])
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}

func thatSliceWith(t *testing.T, _ any) {
	var a any = assert.ThatSliceWith(
		assert_type.NewSoftAssert(t),
		equator.NewComparableEquator[string](),
		[]string{"some val"},
	)
	casted, ok := a.(*asserter.Slice[string])
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
