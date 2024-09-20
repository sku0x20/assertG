package assert

import (
	assertP "github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Bool(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(viaPackage_Bool)
	r.Add(viaVariable_Bool)
	r.Run()
}

func viaPackage_Bool(t *testing.T, e any) {
	var a any = assertP.ThatBool(assertion.NewSoft(t), true)
	casted, ok := a.(*asserter.Bool)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}

func viaVariable_Bool(t *testing.T, e any) {
	assert := assertP.NewAssert(assertion.NewSoft(t))
	var a any = assert.ThatBool(false)
	casted, ok := a.(*asserter.Bool)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
