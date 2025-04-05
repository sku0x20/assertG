package assert

import (
	"github.com/sku0x20/assertG/src/main/assert"
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_String(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(thatString)
	r.Add(thatStringWith)
	r.Run()
}

func thatString(t *testing.T, _ any) {
	var a any = assert.ThatString("some val")
	casted, ok := a.(*asserter.String)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}

func thatStringWith(t *testing.T, _ any) {
	var a any = assert.ThatStringWith(assert_type.NewSoftAssert(t), "some val")
	casted, ok := a.(*asserter.String)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
