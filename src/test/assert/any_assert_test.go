package assert

import (
	assertP "github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/test/assertion"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Any(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(viaPackage)
	r.Run()
}

func viaPackage(t *testing.T, e any) {
	var a any = assertP.ThatAny(assertion.NewFake())
	casted, ok := a.(*asserter.Any)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}

//
//func TestThatAny(t *testing.T) {
//	c := assert.NewCaptureT(t)
//	var a any = c.ThatAny(10)
//	casted, ok := a.(*asserter.AnyAsserter)
//	if casted == nil {
//		t.Fatalf("casted is nil")
//	}
//	if !ok {
//		t.Fatalf("invalid asserter type")
//	}
//}
