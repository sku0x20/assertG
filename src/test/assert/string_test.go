package assert

import (
	assertP "github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/assert_type"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_String(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(viaPackage_String)
	r.Add(viaVariable_String)
	r.Run()
}

func viaPackage_String(t *testing.T, e any) {
	var a any = assertP.ThatString(assert_type.NewSoftAssert(t), "some val")
	casted, ok := a.(*asserter.String)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}

func viaVariable_String(t *testing.T, e any) {
	assert := assertP.NewAssert(assert_type.NewSoftAssert(t))
	var a any = assert.ThatString("some val")
	casted, ok := a.(*asserter.String)
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
