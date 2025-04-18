package assert

import (
	"github.com/sku0x20/assertG/src/main/assert"
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/equator"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Generic(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(that)
	r.Add(thatWith)
	r.Run()
}

func that(t *testing.T, _ any) {
	var a any = assert.That("some val")
	casted, ok := a.(*asserter.Generic[string])
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}

func thatWith(t *testing.T, _ any) {
	var a any = assert.ThatWith(
		assert_type.NewSoftAssert(t),
		equator.NewComparableEquator[string](),
		"some val",
	)
	casted, ok := a.(*asserter.Generic[string])
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
