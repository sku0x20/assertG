package assert

import (
	assertP "github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Any(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(viaPackage_Any)
	r.Add(viaVariable_Any)
	r.Run()
}

func viaPackage_Any(t *testing.T, e any) {
	var a any = assertP.ThatAny(assertion.NewSoft(t), "some val")
	casted, ok := a.(*asserter.Any)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}

func viaVariable_Any(t *testing.T, e any) {
	assert := assertP.NewAssert(assertion.NewSoft(t))
	var a any = assert.ThatAny("some val")
	casted, ok := a.(*asserter.Any)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
