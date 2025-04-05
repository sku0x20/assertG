package assert

import (
	"github.com/sku0x20/assertG/src/main/assert"
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Bool(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(thatBool)
	r.Add(thatBoolWith)
	r.Run()
}

func thatBool(t *testing.T, e any) {
	var a any = assert.ThatBool(true)
	casted, ok := a.(*asserter.Bool)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}

func thatBoolWith(t *testing.T, e any) {
	var a any = assert.ThatBoolWith(assert_type.NewSoftAssert(t), false)
	casted, ok := a.(*asserter.Bool)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
